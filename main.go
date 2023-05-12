package main

import (
	"fmt"
	"net"
)

func main() {

	pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	defer pc.Close()
	buffer := make([]byte, 1024)
	for {
		n, addr, err := pc.ReadFrom(buffer)
		fmt.Printf("recived request from %s \n", addr)
		if err != nil {
			fmt.Println(err)
		}
		pcRequest, err := net.Dial("udp", "1.1.1.1:53")
		if err != nil {
			fmt.Println(err)
		}
		_, err = pcRequest.Write(buffer[:n])
		if err != nil {
			fmt.Println(err)
		}
		buffer := make([]byte, 1024)
		n, err = pcRequest.Read(buffer)
		if err != nil {
			fmt.Println(err)
		}
		_, err = pc.WriteTo(buffer[:n], addr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
