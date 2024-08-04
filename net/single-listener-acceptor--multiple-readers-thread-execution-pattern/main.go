package main

import (
	"fmt"
	"net"
)

const address = "127.0.0.1:8080"

func main() {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Print("Failed creating listener on: ", address)
		return
	}

	defer listener.Close()

	fmt.Print("Listening on: ", address)

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("Failed to accept connection:", err)
			return
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		// Read data from the connection
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		request := string(buffer[:n])
		fmt.Println("Received:", request)

		// Write the same data back to the connection (echo)
		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}
