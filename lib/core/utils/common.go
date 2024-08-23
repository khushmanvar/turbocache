package utils

import (
	"fmt"
	"strconv"
)

func StrToInt(str string) int64 {
	intVal, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		WriteErrToConsole("fail to convert string to int")
	}
	return intVal
}

func WriteErrToConsole(str string) {
	fmt.Errorf("Error: ", str)
}
