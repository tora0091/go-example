package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:7777")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	f, err := os.Open("a.jpeg")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	n, err := conn.Write([]byte(b))
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("[client ] send size: %d", n)
}
