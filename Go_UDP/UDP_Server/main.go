package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	//Listening to incoming udp packets

	udpserver, err := net.ListenPacket("udp", ":1053") /*listens on all available IP addresses of
	  the local system except multicast IP addresses.*/
	if err != nil {
		log.Fatal(err)
	}

	defer udpserver.Close() // Ensures the UDP connection is closed when the main function exits.

	for {

		buf := make([]byte, 1024)
		_, addr, err := udpserver.ReadFrom(buf) //getting data from the UDP client.
		if err != nil {
			continue //Since UDP, Packet loss doesn't matter
		}

		go response(udpserver, addr, buf)
	}
}

func response(udpserver net.PacketConn, addr net.Addr, buf []byte) {

	receivedTime := time.Now().Format(time.ANSIC)
	fmt.Printf("Received message from (%v): %v!. at %v", addr, string(buf), receivedTime)
	responsestr1 := fmt.Sprintf("your message: %v.", string(buf))
	responsestr2 := fmt.Sprintf("Received by the server at:%v", receivedTime)
	udpserver.WriteTo([]byte(responsestr1), addr) //Sends the response string back to the client.
	udpserver.WriteTo([]byte(responsestr2), addr)
}
