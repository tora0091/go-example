package main

import (
	"log"
	"net"
	"os/exec"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":15051")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go func() {
			defer conn.Close()
			for {
				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				if err != nil {
					log.Println(err)
					return
				}
				log.Println(string(buf[:n]))

				command, options := getCommands(string(buf[:n]))
				out, err := exec.Command(command, options...).Output()
				if err != nil {
					log.Println(err)
					return
				}

				_, err = conn.Write(out)
				if err != nil {
					log.Println(err)
					return
				}
			}
		}()
	}
}

func getCommands(line string) (string, []string) {
	list := strings.Split(line, " ")

	command := list[0]
	var options []string
	if len(list) > 1 {
		options = list[1:]
	}
	return command, options
}
