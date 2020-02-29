package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("[server] starting server: localhost:8888")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()

		go func() {
			fmt.Printf("[server] remote address: %v\n", conn.RemoteAddr())

			err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
			if err != nil {
				log.Fatalln(err)
			}

			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatalln(err)
			}

			msg := string(buf[:n])

			fmt.Printf("[server] %s\n", msg)
		}()
	}
}
