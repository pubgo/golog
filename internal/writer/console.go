package zwriter

import (
	"github.com/rs/zerolog"
	"os"
)

func NewConsole() *console {
	return &console{c: &zerolog.ConsoleWriter{Out: os.Stderr}}
}

type console struct {
	c *zerolog.ConsoleWriter
}

func (t *console) Write(p []byte) (n int, err error) {
	return t.c.Write(p)
}
