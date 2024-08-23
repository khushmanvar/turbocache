package cmd

import (
	"errors"
	"fmt"
	"io"
	"turbocache/lib/core/store"
	"turbocache/lib/core/types"
	"turbocache/lib/core/utils"
)

func evalPING(args []string, c io.ReadWriter) error {
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

func evalSET(args []string, c io.ReadWriter) *types.Exception {
	if len(args) < 3 {
		return types.NewException("wrong number of arguments")
	}

	var key, value string

	key, value = args[0], args[1]

	store.Put(key, store.NewRecord(value, utils.GetExpiresAt(utils.StrToInt(args[2]))))
	c.Write([]byte("+OK\r\n"))

	return nil
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

func EvalAndRespond(cmd *types.TurboCommand, c io.ReadWriter) error {
	switch cmd.Cmd {
	case "PING":
		return evalPING(cmd.Args, c)
	default:
		return evalPING(cmd.Args, c)
	}
}
