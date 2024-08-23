package utils

import (
	"bufio"
	"errors"
	"io"
	"net"
	"strings"
	"turbocache/lib/core/types"
)

func ReadCommand(c io.ReadWriter) (*types.TurboCommand, error) {
	// TODO: Max read in one shot is 512 bytes
	// To allow input > 512 bytes, then repeated read until
	// we get EOF or designated delimiter
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return nil, err
	}

	tokens, err := DecodeArrayString(buf[:n])
	if err != nil {
		return nil, err
	}

	return &types.TurboCommand{
		Cmd:  strings.ToUpper(tokens[0]),
		Args: tokens[1:],
	}, nil
}

func ReadCommandFromConn(conn net.Conn) (*types.TurboCommand, error) {
	// Create a buffered reader to read from the connection
	reader := bufio.NewReader(conn)

	// Read a line from the connection (up to newline)
	cmdStr, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Trim any trailing newline or carriage return characters
	cmdStr = strings.TrimSpace(cmdStr)

	// Validate the command (optional, depends on your use case)
	if cmdStr == "" {
		return nil, errors.New("received an empty command")
	}

	// Parse the command string into a TurboCommand
	parts := strings.Fields(cmdStr)
	if len(parts) == 0 {
		return nil, errors.New("invalid command format")
	}

	command := parts[0]
	args := parts[1:]

	turboCommand := &types.TurboCommand{
		Cmd:  command,
		Args: args,
	}

	return turboCommand, nil
}

func Respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
