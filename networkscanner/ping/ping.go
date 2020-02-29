package ping

import (
	"net"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func Ping(addr string) (string, bool, error) {
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return addr, false, err
	}
	defer c.Close()

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("HELLO-R-U-THERE"),
		},
	}

	wb, err := wm.Marshal(nil)
	if err != nil {
		return addr, false, err
	}

	if _, err = c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(addr)}); err != nil {
		return addr, false, err
	}

	rb := make([]byte, 1500)
	n, _, err := c.ReadFrom(rb)
	if err != nil {
		return addr, false, err
	}

	rm, err := icmp.ParseMessage(1, rb[:n])
	if err != nil {
		return addr, false, err
	}

	switch rm.Type {
	case ipv4.ICMPTypeEchoReply:
		return addr, true, nil
	default:
		return addr, false, nil // unreachable
	}
}
