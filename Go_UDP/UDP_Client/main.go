package main

import (
	"fmt"
	"net"
	"os"
)

/*type player struct{

	name string



} */

func main() {

	udpserver, err := net.ResolveUDPAddr("udp", ":1053")
	if err != nil {
		println("ResolveUDPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialUDP("udp", nil, udpserver)
	if err != nil {
		println("Listening UDP failed:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	_, err = conn.Write([]byte("Hello UDP Server"))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	received1 := make([]byte, 1024)
	received2 := make([]byte, 1024)
	n, _, err := conn.ReadFrom(received1)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(received1[:n]))
	p, _, err := conn.ReadFrom(received2)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(received2[:p]))
}
