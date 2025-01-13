package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	clients    = make(map[net.Conn]string)
	clientsMux sync.Mutex // Mutex to handle concurrent access to the clients
)

func broadcastMessage(sender net.Conn, message string) {

	clientsMux.Lock()
	defer clientsMux.Unlock()

	for client := range clients {
		if client != sender { // Don't send the message back to the sender
			_, err := client.Write([]byte("Broadcast: " + message + "\n"))
			if err != nil {
				fmt.Printf("Error sending message to client: %v\n", err)
			}
		}
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		clientsMux.Lock()
		delete(clients, conn) // Remove the client from the active list
		clientsMux.Unlock()
		conn.Close()
		fmt.Printf("Client disconnected: %s\n", conn.RemoteAddr().String())
	}()

	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Client connected: %s\n", clientAddr)
	// critical section
	clientsMux.Lock()
	clients[conn] = clientAddr
	clientsMux.Unlock()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Message from %s: %s\n", clientAddr, message)

		if strings.TrimSpace(message) == "exit" {
			fmt.Printf("Client %s disconnected.\n", clientAddr)
			return
		}

		// Broadcast the message to all clients except the sender
		broadcastMessage(conn, message)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from client: %v\n", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8008")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server started on port 8008")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
