package cmd

import (
	"fmt"
	"io"
	"turbocache/lib/core/store"
	"turbocache/lib/core/types"
	"turbocache/lib/core/utils"
)

var RESP_NIL []byte = []byte("$-1\r\n")

func evalPING(args []string, c io.ReadWriter) *types.Exception {
	var b []byte

	if len(args) >= 2 {
		return types.NewException("ERR wrong number of arguments for 'ping' command")
	}

	if len(args) == 0 {
		b = encode("PONG", true)
	} else {
		b = encode(args[0], false)
	}

	_, err := c.Write(b)
	if err != nil {
		utils.WriteErrToConsole("Fail to write")
	}

	return nil
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

func evalGET(args []string, c io.ReadWriter) *types.Exception {
	if len(args) < 2 {
		return types.NewException("invalid args")
	}

	var key = args[0]
	var rcd *types.Record

	rcd = store.Get(key)
	c.Write(encode(rcd.Value, false))

	return nil
}

func encode(value interface{}, isSimple bool) []byte {
	switch v := value.(type) {
	case string:
		if isSimple {
			return []byte(fmt.Sprintf("+%s\r\n", v))
		}
		return []byte(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v))
	case int64:
		return []byte(fmt.Sprintf(":%d\r\n", v))
	default:
		return RESP_NIL
	}
}

func EvalAndRespond(cmd *types.TurboCommand, c io.ReadWriter) *types.Exception {
	switch cmd.Cmd {
	case "PING":
		return evalPING(cmd.Args, c)
	case "GET":
		return evalGET(cmd.Args, c)
	case "SET":
		return evalSET(cmd.Args, c)
	default:
		return evalPING(cmd.Args, c)
	}
}
