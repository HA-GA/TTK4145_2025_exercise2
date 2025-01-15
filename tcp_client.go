package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ServerAddr, _ := net.ResolveTCPAddr("tcp", "10.100.23.204:33546")
	ClientAddr, _ := net.ResolveTCPAddr("tcp", "10.100.23.37:30027")

	conn, _ := net.DialTCP("tcp", nil, ServerAddr)
	defer conn.Close()

	listener, _ := net.ListenTCP("tcp", ClientAddr)
	defer listener.Close()

	go sendConnection(conn)
	go acceptConnection(listener)

	select {}

}

func sendConnection(conn *net.TCPConn) {
	fmt.Println("Server has connected")

	buffer := make([]byte, 1024)
	message := "Connect to: 10.100.23.37:30027"
	copy(buffer, message)
	_, err := conn.Write([]byte(buffer))
	if err != nil {
		fmt.Println("Error with sending message")
		return
	}

	buffer2 := make([]byte, 1024)
	n, err := conn.Read(buffer2)
	if err != nil {
		fmt.Println("Error with reading message")
		return
	}

	fmt.Printf("Received: ", string(buffer2[:n]))

	fmt.Println("Message sendt to server", message)
}

func acceptConnection(listener *net.TCPListener) {

	accept, err := listener.Accept()

	if err != nil {
		log.Fatal(err)
		return
	}
	defer accept.Close()
	acceptMessage := make([]byte, 1024)

	_, err2 := accept.Read(acceptMessage)
	if err2 != nil {
		fmt.Printf("Error with receiving accepte message")
		return
	}

	fmt.Printf("accept message: ", string(acceptMessage))

	buffer := make([]byte, 1024)
	message := "\n\nhei hei\000"
	copy(buffer, message)

	_, err = accept.Write([]byte(buffer))
	if err != nil {
		fmt.Println("Error with sending message")
		return
	}

	buffer2 := make([]byte, 1024)
	n, err := accept.Read(buffer2)
	if err != nil {
		fmt.Println("Error with reading message")
		return
	}
	fmt.Printf("Received: ", string(buffer2[:n]))

}
