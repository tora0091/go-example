package main

import (
	"fmt"
	"log"

	whois "github.com/brimstone/golang-whois"
)

func main() {
	w, err := whois.GetWhois("google.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(w)
}
