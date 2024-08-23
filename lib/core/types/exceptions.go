package types

import "errors"

type Exception struct {
	err error
}

func NewException(err string) *Exception {
	exp := &Exception{errors.New(err)}
	return exp
}
