package main

import (
	"log"
	"net"

	"go-example/packetcapture/message"
)

func init() {
	log.SetPrefix("[client] ")
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	log.Printf("ok, %s\n", conn.RemoteAddr())

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(buf[:n]))

	// read GOB data
	buf = make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	target := message.NewTarget()
	err = target.Unmarshal(buf[:n])
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(target)
}
