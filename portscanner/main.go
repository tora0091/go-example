package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
)

type Status struct {
	Port   int
	Result bool
}

func main() {
	hostname := flag.String("hostname", "localhost", "scanning hostname or address")
	protocol := flag.String("protocol", "tcp", "scanning protocol type, tcp or udp")
	flag.Parse()

	fmt.Printf("Hostname: %s, Protocol: %s\n", *hostname, *protocol)
	for port := 0; port < 1024; port++ {
		status := Scanning(*hostname, *protocol, port)
		ShowResult(status)
	}
}

func Scanning(hostname, protocol string, port int) Status {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 10*time.Second)
	if err != nil {
		return Status{Port: port, Result: false}
	}
	defer conn.Close()
	return Status{Port: port, Result: true}
}

func ShowResult(status Status) {
	if status.Result == true {
		fmt.Printf("  Port: %3d, Result: %t\n", status.Port, status.Result)
	}
}
