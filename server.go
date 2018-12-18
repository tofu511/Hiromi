package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

const (
	Host string = "localhost"
	Port string = "5163"
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

			fmt.Fprint(conn, "HTTP/1.1 240 Exotic Japan!\r\n\r\n Hi!")
			conn.Close()
		}()
	}
}
