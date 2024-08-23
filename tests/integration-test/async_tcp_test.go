package main

import (
	"fmt"
	"net"
	"sync"
)

const (
	serverAddress = "127.0.0.1:7379" // Replace with your server address
	numClients    = 1000000          // Number of concurrent clients
)

func main() {
	var wg sync.WaitGroup
	wg.Add(numClients)

	for i := 0; i < numClients; i++ {
		go func(clientID int) {
			defer wg.Done()

			conn, err := net.Dial("tcp", serverAddress)
			if err != nil {
				fmt.Printf("Client %d: Error connecting to server: %v\n", clientID, err)
				return
			}
			defer conn.Close()

			// Send a message
			message := fmt.Sprintf("Hello from client %d\n", clientID)
			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Printf("Client %d: Error sending data: %v\n", clientID, err)
				return
			}

			// Optionally, read response from the server
			buffer := make([]byte, 1024)
			_, err = conn.Read(buffer)
			if err != nil {
				fmt.Printf("Client %d: Error reading response: %v\n", clientID, err)
				return
			}
		}(i)
	}

	// Wait for all clients to finish
	wg.Wait()
	fmt.Println("All clients have finished.")
}
