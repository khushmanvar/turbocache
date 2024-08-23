package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"turbocache/config"
	"turbocache/lib/core/types"
	"turbocache/lib/core/utils"
)

func RunSyncTCPServer() {
	log.Println("starting a synchronous TCP server on", config.Host, config.Port)

	var con_clients int = 0

	// listening to the configured host:port
	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		panic(err)
	}

	for {
		// blocking call: waiting for the new client to connect
		c, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}

		// increment the number of concurrent clients
		con_clients += 1
		log.Println("client connected with address:", c.RemoteAddr(), "concurrent clients", con_clients)

		for {
			// over the socket, continuously read the command and print it out
			cmd, err := utils.ReadCommand(c)
			if err != nil {
				c.Close()
				con_clients -= 1
				log.Println("client disconnected", c.RemoteAddr(), "concurrent clients", con_clients)
				if err == io.EOF {
					break
				}
				log.Println("err", err)
			}
			respond(cmd, c)
		}
	}
}

func respondError(err error, c net.Conn) {
	c.Write([]byte(fmt.Sprintf("-%s\r\n", err)))
}

func respond(cmd *types.TurboCommand, c net.Conn) {
	err := utils.EvalAndRespond(cmd, c)
	if err != nil {
		respondError(err, c)
	}
}