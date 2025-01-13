package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client connected: %s\n", clientAddr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		message := scanner.Text()
		fmt.Printf("Message from %s: %s\n", clientAddr, message)

		if strings.TrimSpace(message) == "exit" {
			fmt.Printf("Client %s disconnected.\n", clientAddr)
			return
		}

		conn.Write([]byte("Server: " + message + "\n"))

	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from client: %v\n", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8008")
	if err != nil {
		fmt.Println("Error Starting Server", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server Started")
	// start Accepting the server

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting Connection", err)
			continue
		}
		go handleConnection(conn)
	}

}
