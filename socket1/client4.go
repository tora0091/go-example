package main

import (
	"log"
	"net"
)

func init() {
	log.SetPrefix("[client] ")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	file := "/tmp/server-socket-test.sock"

	conn, err := net.Dial("unix", file)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("abcdefghijklmnopqrstuvwxyz01234567890HELLO!!"))
	if err != nil {
		log.Fatalln(err)
	}

	err = conn.(*net.UnixConn).CloseWrite()
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
}
