package main

import (
	"fmt"
	"net"
	"sync"
)

func scan(host string, port int, wg *sync.WaitGroup) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	defer conn.Close()
	defer wg.Done()
	fmt.Println("[+] connection established.", conn.RemoteAddr())
}

func main() {
	host := "scanning host name or ip address"

	var wg sync.WaitGroup
	for port := 1; port <= 100; port++ {
		wg.Add(1)
		go scan(host, port, &wg)
	}
	wg.Wait()
}
