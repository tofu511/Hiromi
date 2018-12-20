package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

const (
	Host         string = "localhost"
	Port         string = "5163"
	DocumentRoot        = "./public/"
)

func main()  {
	endpoint := strings.Join([]string{Host, Port}, ":")
	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Server is running at %s\n", endpoint)

	for {
		// 接続するまで待つ
		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// クライアントからのリクエストをパースする
			request, err := http.ReadRequest(bufio.NewReader(conn))

			if err != nil {
				panic(err)
			}

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(dump))

			file, err := ioutil.ReadFile(DocumentRoot + "index.html")

			if err != nil {
				panic(err)
			}

			response := createResponse("240 Exotic Japan!", string(file))

			fmt.Fprint(conn, response)
			
			conn.Close()
		}()
	}
}

func createResponse(statusCode, body string) string {
	now := time.Now().Format(time.RFC1123)
	server := "Hiromi"
	return fmt.Sprintf(`HTTP/1.1 %s 
Server: %s
Date: %s
Connection: Close

%s`, statusCode, server, now, body)
}
