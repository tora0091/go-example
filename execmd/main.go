package main

import (
	"flag"
	"log"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

const letter = "abcdef1234567890"

var iface string
var mac string

func init() {
	flag.StringVar(&iface, "i", "eth0", "target interface (ex: eth0)")
	flag.StringVar(&mac, "m", "", "you want to set mac address")
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()

	// ifconfig eth0
	ifconfigCmd := exec.Command("ifconfig", iface)
	if err := ifconfigCmd.Run(); err != nil {
		log.Fatalf("error, %s is not found\n", iface)
	}

	// ifconfig eth0 down
	if err := exec.Command("ifconfig", iface+" down").Run(); err != nil {
		log.Fatalf("error, ifconfig %s down\n", iface)
	}

	if mac == "" {
		mac = randMacAddress()
	}

	// ifconfig eth0 hw ether [mac address]
	if err := exec.Command("ifconfig", iface+" hw ether "+mac).Run(); err != nil {
		log.Fatalf("error, ifconfig %s hw ether %s\n", iface, mac)
	}

	// ifconfig eth0 up
	if err := exec.Command("ifconfig", iface+" up").Run(); err != nil {
		log.Fatalf("error, ifconfig %s up\n", iface)
	}

	ifconfigCmd.Run()
}

func randMacAddress() string {
	mac := []string{}
	for i := 0; i < 6; i++ {
		mac = append(mac, randOctet())
	}
	return strings.Join(mac, ":")
}

func randOctet() string {
	buf := make([]byte, 2)
	for i := range buf {
		buf[i] = letter[rand.Intn(len(letter))]
	}
	return string(buf)
}
