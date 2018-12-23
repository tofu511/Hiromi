package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	Host         string = "localhost"
	Port         string = "5163"
	DocumentRoot        = "./public"
)

var contentTypeMap = map[string]string{
	".html": "text/html",
	".htm":  "text/html",
	".txt":  "text/plain",
	".css":  "text/css",
	".png":  "image/png",
	".jpg":  "image/jpg",
	".jpeg": "image/jpeg",
	".gif":  "image/gif",
	".ico":  "image/x-icon",
}

func main() {
	endpoint := strings.Join([]string{Host, Port}, ":")
	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go func() {
			request := parseRequest(conn)

			filePath := convertPath(DocumentRoot + request.URL.Path)
			file := readFileFromUrlPath(filePath)

			lang := request.Header.Get("Accept-Language")
			status := createStatus(filePath, lang)

			contentType, _ := contentTypeMap[path.Ext(filePath)]

			response := createResponse(status, contentType, string(file))

			fmt.Fprint(conn, response)

			conn.Close()
		}()
	}
}

func createStatus(path, acceptLang string) string {
	statusCode := 200
	statusText := "OK"
	if !exists(path) {
		statusCode = 404
		statusText = "Not Found"
	}

	if statusCode == 200 && strings.Split(acceptLang, ",")[0] == "ja-JP" {
		statusCode = 240
		statusText = "Exotic Japan!"
	}

	return strings.Join([]string{strconv.Itoa(statusCode), statusText}, " ")
}

func createResponse(statusCode, contentType, body string) string {
	now := time.Now().Format(time.RFC1123)
	server := "Hiromi"
	return fmt.Sprintf(`HTTP/1.1 %s 
Server: %s
Date: %s
Connection: close
Content-type: %s

%s`, statusCode, server, now, contentType, body)
}

func parseRequest(reader io.Reader) *http.Request {
	request, err := http.ReadRequest(bufio.NewReader(reader))

	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	return request
}

func readFileFromUrlPath(path string) []byte {
	if !exists(path) {
		path = DocumentRoot + "/404.html"
	}

	file, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	return file
}

func convertPath(path string) string {
	if path == DocumentRoot + "/" {
		path = DocumentRoot + "/index.html"
	}
	return path
}

func exists(filePath string) bool {
	_, err := os.Stat(filePath)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	return err == nil
}