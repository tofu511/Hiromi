package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

const (
	Host         string = "localhost"
	Port         string = "5163"
	DocumentRoot        = "./public"
)

func main()  {
	endpoint := strings.Join([]string{Host, Port}, ":")
	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		panic(err)
	}

	for {
		// 接続するまで待つ
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go func() {
			request := parseRequest(conn)

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(dump))

			path := convertPath(DocumentRoot + request.URL.Path)
			file := readFileFromUrlPath(path)

			status := "200 OK"
			if !exists(path) {
				status = "404 Not Found"
			}

			response := createResponse(status, "text/html", string(file))

			fmt.Fprint(conn, response)

			conn.Close()
		}()
	}
}

func createResponse(statusCode, contentType, body string) string {
	now := time.Now().Format(time.RFC1123)
	server := "Hiromi"
	return fmt.Sprintf(`HTTP/1.1 %s 
Server: %s
Date: %s
Connection: Close
Content-type: %s

%s`, statusCode, server, now, contentType, body)
}

func parseRequest(reader io.Reader) *http.Request {
	request, err := http.ReadRequest(bufio.NewReader(reader))

	if err != nil {
		panic(err)
	}
	return request
}

func readFileFromUrlPath(path string) []byte {
	if !exists(path) {
		path = DocumentRoot + "/404.html"
	}

	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return file
}

func convertPath(path string) string {
	if path == "/" {
		path = "/index.html"
	}
	return path
}

func exists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}