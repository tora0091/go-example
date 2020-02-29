package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Println(err)
	}

	request, err := http.NewRequest("GET", "http://localhost:9999", nil)
	if err != nil {
		log.Println(err)
	}
	request.Write(conn)

	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		log.Println(err)
	}

	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(dump))
}
