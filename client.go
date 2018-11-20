package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main()  {
	host := "localhost"
	port := "5163"
	endpoint := strings.Join([]string{host, port}, ":")
	conn, err := net.Dial("tcp", endpoint)

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(http.MethodGet, "http://" + endpoint, nil)

	if err != nil {
		panic(err)
	}

	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)

	if err != nil {
		panic(err)
	}

	dump, err := httputil.DumpResponse(response, true)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dump))

}
