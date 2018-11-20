package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)


func main()  {
	host := "localhost"
	port := "5163"
	endpoint := strings.Join([]string{host, port}, ":")
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

			response := http.Response{
				StatusCode:200,
				ProtoMajor:1,
				ProtoMinor:0,
				Body: ioutil.NopCloser(
					strings.NewReader("Exotic Japan!!\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
