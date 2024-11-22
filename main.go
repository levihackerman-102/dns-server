package main

import (
	"fmt"
	"net"

	"dns-server/pkg/dns"
)

func main() {
	fmt.Printf("Starting DNS Server...\n")
	
	packetConnection, err := net.ListenPacket("udp", ":5353")
	if err != nil {
		panic(err)
	}
	defer packetConnection.Close()

	for {
		buf := make([]byte, 512)
		bytesRead, addr, err := packetConnection.ReadFrom(buf)
		if err != nil {
			fmt.Printf("Read error from %s: %s\n", addr.String(), err)
			continue
		}
		go dns.HandlePacket(packetConnection, addr, buf[:bytesRead])
	}
}
