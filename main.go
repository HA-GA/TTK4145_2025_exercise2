package main

import (
	"fmt"
	"net"
)

func main() {
	ServerAddr, _ := net.ResolveUDPAddr("udp", ":20027")
	conn, _ := net.ListenUDP("udp", ServerAddr)
	defer conn.Close()

	conn, err := net.DialUDP("udp", nil, ServerAddr)
	if err != nil {
		fmt.Println("Kunne ikke opprette tilkobling")
	}
	defer conn.Close()

	message := "Hei fra gr44"
	_, err = conn.Write([]byte(message))

	fmt.Println("Melding sendt til server")
}
