package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":15051")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			defer conn.Close()
			for {
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					// log.Fatalln(err)
					return
				}
				log.Println(string(buf[:n]))
			}
		}()
	}
}
