package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	hostname := flag.String("h", "google.com", "look up hostname")
	flag.Parse()

	addrs, err := net.LookupHost(*hostname)
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
