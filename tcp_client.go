package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    serverAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:33546")
    conn, _ := net.DialTCP("tcp", nil, serverAddr)
    defer conn.Close()

    fmt.Println("Connected to server:", serverAddr.String())

    reader := bufio.NewReader(conn)
    greeting, _ := reader.ReadString('\n')
    fmt.Print("Server says: ", greeting)

    consoleReader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Enter message (type 'quit' to exit): ")
        input, _ := consoleReader.ReadString('\n')
        if input == "quit\n" {
            fmt.Println("Closing connection.")
            return
        }


        conn.Write([]byte(input))

        response, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Server closed the connection.")
            return
        }
        fmt.Print("Server response: ", response)
    }
}
