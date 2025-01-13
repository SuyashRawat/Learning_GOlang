package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8008")
	if err != nil {
		fmt.Println("Error connecting to Server:", err)
		return
	}
	defer conn.Close()

	// Channel to signal when to exit
	exitChan := make(chan bool)

	// Goroutine to handle receiving messages from the server
	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println("Error reading from Server:", err)
				exitChan <- true
				return
			}
			fmt.Print(message) // Print broadcasted messages from the server
		}
	}()

	// Main loop to send messages to the server
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error writing to Server:", err)
			return
		}

		if message == "exit" {
			exitChan <- true
			break
		}
	}

	<-exitChan // Wait for the exit signal before closing the client
}
