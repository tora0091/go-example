package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	host := flag.String("host", "localhost", "set target host name or ip address")
	flag.Parse()

	hostname := *host + ":15051"
	conn, err := net.Dial("tcp", hostname)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}

		_, err = conn.Write(line)
		if err != nil {
			log.Fatalln(err)
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(buf[:n]))
	}
}
