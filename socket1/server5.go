package main

import (
	"io"
	"log"
	"net"
	"os"
)

func init() {
	log.SetPrefix("[server] ")
}
func main() {
	listen, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()

		log.Println(conn.RemoteAddr())

		data := make([]byte, 0)
		go func() {
			for {
				buf := make([]byte, 1024)

				n, err := conn.Read(buf)
				if err != nil {
					if err != io.EOF {
						log.Fatalln(err)
					}
					break
				}

				buf = buf[:n]
				data = append(data, buf...)
			}
			log.Printf("recive size: %d\n", len(data))

			f, err := os.Create("b.jpeg")
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close()

			_, err = f.Write(data)
			if err != nil {
				log.Fatalln(err)
			}
		}()
	}
}
