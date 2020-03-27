package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
	device  string        = "wlp3s0"
	snaplen int32         = 1024
	promisc bool          = false
	timeout time.Duration = pcap.BlockForever
	filter  string        = "tcp and port 443"
)

func main() {
	f, err := os.Create("test.pcap")
	if err != nil {
		log.Fatalln(err)
	}

	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snaplen), layers.LinkTypeEthernet)
	defer f.Close()

	handle, err := pcap.OpenLive(device, snaplen, promisc, timeout)
	if err != nil {
		log.Fatalln(err)
	}
	defer handle.Close()

	packetCount := 0
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
		w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		packetCount++

		if packetCount > 100 {
			break
		}
	}
}
