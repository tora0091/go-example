package main

import (
	"io"
	"log"
	"net"
	"os"
)

func init() {
	log.SetPrefix("[server] ")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	file := "/tmp/server-socket-test.sock"
	defer os.Remove(file)

	listen, err := net.Listen("unix", file)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			defer conn.Close()
			buf := make([]byte, 10)

			n, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					log.Fatalln(err)
				}
			}

			buf = buf[:n]
			log.Printf("receive: %v\n", buf)
		}()
	}
}
