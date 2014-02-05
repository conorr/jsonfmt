package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

const READBYTES int = 1024

func ReadFile(filename string) *bytes.Buffer {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var buf bytes.Buffer
	data := make([]byte, READBYTES)
	for {
		n, err := fi.Read(data)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			os.Exit(1)
		}
		if n == 0 {
			break
		}
		buf.Write(data[:n])
	}
	fi.Close()
	return &buf
}

func WriteFile(filename string, buf *bytes.Buffer) error {
	fo, err := os.Create(filename)
	if err != nil {
		return err
	}
	fo.Write(buf.Bytes())
	fo.Close()
	return nil
}
