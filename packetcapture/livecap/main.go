// ex: tcpdump -X -i wlp3s0 tcp and port 443
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	device  string        = "wlp3s0"
	snaplen int32         = 1024
	promisc bool          = false
	timeout time.Duration = pcap.BlockForever
	filter  string        = "tcp and port 443"
)

func main() {
	handle, err := pcap.OpenLive(device, snaplen, promisc, timeout)
	if err != nil {
		log.Fatalln(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatalln(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// applayer := packet.ApplicationLayer()
		// if applayer == nil {
		// 	continue
		// }
		// payload := applayer.Payload()
		// fmt.Println(string(payload))

		fmt.Printf("%s\n", packet)
		fmt.Printf("%s", hex.Dump(packet.Data()))
	}
}
