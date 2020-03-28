package main

import (
	"log"
	"net"

	"go-example/packetcapture/message"
)

func init() {
	log.SetPrefix("[server] ")
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	log.Printf("ok, %s\n", conn.RemoteAddr())

	n, err := conn.Write([]byte("hello " + conn.RemoteAddr().String() + ", I am server. ok!!"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send data %d bytes to %s\n", n, conn.RemoteAddr())

	// send GOB data
	target := message.NewTarget()
	target.Name = "Yashushi Tomita"
	target.Age = 58
	target.Area = "Shibuya in Tokyo, Japan"
	target.Job = "Diplomat"
	data, err := target.Marshal()
	if err != nil {
		log.Fatalln(err)
	}
	n, err = conn.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("send GOB data %d bytes to %s\n", n, conn.RemoteAddr())
}
