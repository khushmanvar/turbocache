package types

import (
	"syscall"
)

type TurboCommand struct {
	Cmd  string
	Args []string
}

type FDCommand struct {
	Fd int
}

func (f FDCommand) Write(b []byte) (int, error) {
	return syscall.Write(syscall.Handle(f.Fd), b)
}

func (f FDCommand) Read(b []byte) (int, error) {
	return syscall.Read(syscall.Handle(f.Fd), b)
}
