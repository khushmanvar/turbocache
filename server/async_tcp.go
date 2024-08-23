package server

import (
	"log"
	"net"
	"strconv"
	"sync"
	"turbocache/config"
	"turbocache/lib/core/utils"
)

var conClients int = 0
var mu sync.Mutex // To ensure thread-safe operations on conClients

func RunAsyncTCPServer() error {
	log.Println("starting an asynchronous TCP server on", config.Host, config.Port)

	// Convert config.Port (int) to string
	portStr := strconv.Itoa(config.Port)

	listener, err := net.Listen("tcp", net.JoinHostPort(config.Host, portStr))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle each connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Get the client address and log it
	clientAddr := conn.RemoteAddr().String()
	log.Printf("Client connected: %s", clientAddr)

	mu.Lock()
	conClients++
	mu.Unlock()

	// Instead of using FDCommand, read directly from the connection
	cmd, err := utils.ReadCommandFromConn(conn)
	if err != nil {
		log.Println("Error reading command:", err)
		mu.Lock()
		conClients--
		mu.Unlock()
		return
	}

	// Respond to the command
	respond(cmd, conn)

	mu.Lock()
	conClients--
	mu.Unlock()
}
