package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[server] starting server : localhost:9999")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}

		go func() {
			fmt.Printf("[server] remote address : %v\n", conn.RemoteAddr())

			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				log.Println(err)
			}

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(dump))

			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello World!!")),
			}

			response.Write(conn)
			conn.Close()
		}()
	}
}
