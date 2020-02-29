package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"networkscanner/ping"
)

type SEARCH_TYPE string

var search_all SEARCH_TYPE = "all"
var search_single SEARCH_TYPE = "single"

func main() {
	ipaddr := flag.String("i", "192.168.0.1", "ping ip address")
	flag.Parse()

	search, err := checkIpAddr(*ipaddr)
	if err != nil {
		log.Fatal(err)
	}

	switch search {
	case search_single:
		Scan(*ipaddr)
	case search_all:
		for i := 0; i <= 255; i++ {
			// todo: goroutine
			Scan(*ipaddr + "." + strconv.Itoa(i))
		}
	}
}

func checkIpAddr(ipaddr string) (SEARCH_TYPE, error) {
	addr := strings.Split(ipaddr, ".")
	for _, v := range addr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return "", err
		}
		if num < 0 || num > 255 {
			return "", fmt.Errorf("Error: range error, from 0 to 255")
		}
	}

	var s SEARCH_TYPE
	switch len(addr) {
	case 3:
		s = search_all
	case 4:
		s = search_single
	default:
		return "", fmt.Errorf("Error: format error, ex: 192.168.0 or 192.168.0.1")
	}
	return s, nil
}

func Scan(ipaddr string) {
	address, result, err := ping.Ping(ipaddr)
	if err != nil {
		log.Fatal(err)
	}
	if result {
		fmt.Printf("%s %t\n", address, result)
	}
}
