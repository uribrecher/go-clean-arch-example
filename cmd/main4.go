package main

import (
	"clean/internal"
	"io"
	"os"
)

func encrypt2(reader io.Reader, writer io.Writer) error {
	var bytes []byte
	if _, err := reader.Read(bytes); err != nil {
		return err
	}
	result := internal.Translate(bytes)
	if _, err := writer.Write(result); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := encrypt2(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}
