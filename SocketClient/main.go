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
		fmt.Println("Error connecting to Server", err)
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)

		message, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(message))
		message = strings.TrimSpace(message)
		if string(message) == "exit" {
			// fmt.Println(message)
			break
		}
		if err != nil {
			fmt.Println("Error writing to Server", err)
			return
		}

	}
	conn.Close()
}
