package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		conn.Write([]byte(scanner.Text()))
	}
}
