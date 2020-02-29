package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fr, err := os.Open("010.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer fr.Close()

	b, err := ioutil.ReadAll(fr)
	if err != nil {
		log.Fatalln(err)
	}

	fw, err := os.Create("020.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer fw.Close()

	_, err = fw.Write(b)
	if err != nil {
		log.Fatalln(err)
	}
}
