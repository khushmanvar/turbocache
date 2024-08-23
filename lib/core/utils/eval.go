package utils

import (
	"errors"
	"fmt"
	"net"
	"turbocache/lib/core/types"
)

func evalPING(args []string, c net.Conn) error {
	var b []byte

	if len(args) >= 2 {
		return errors.New("ERR wrong number of arguments for 'ping' command")
	}

	if len(args) == 0 {
		b = encode("PONG", true)
	} else {
		b = encode(args[0], false)
	}

	_, err := c.Write(b)
	return err
}

func encode(value interface{}, isSimple bool) []byte {
	switch v := value.(type) {
	case string:
		if isSimple {
			return []byte(fmt.Sprintf("+%s\r\n", v))
		}
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v))
	}
	return []byte{}
}

func EvalAndRespond(cmd *types.TurboCommand, c net.Conn) error {
	switch cmd.Cmd {
	case "PING":
		return evalPING(cmd.Args, c)
	default:
		return evalPING(cmd.Args, c)
	}
}