package utils

import (
	"net"
	"strings"
	"turbocache/lib/core/types"
)

func ReadCommand(c net.Conn) (*types.TurboCommand, error) {
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

func Respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}