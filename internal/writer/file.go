package zwriter

import (
	"os"
)

func NewFileWriter(path string) *fileWriter {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err.Error())
	}
	return &fileWriter{f: f}
}

type fileWriter struct {
	f *os.File
}

func (t *fileWriter) Write(p []byte) (n int, err error) {
	return t.f.Write(p)
}

func (t *fileWriter) Close() error {
	return t.f.Close()
}
