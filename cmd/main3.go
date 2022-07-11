package main

import (
	"clean/internal"
	"os"
)

func encrypt() error {
	var bytes []byte
	if _, err := os.Stdin.Read(bytes); err != nil {
		return err
	}
	result := internal.Translate(bytes)
	if _, err := os.Stdout.Write(result); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := encrypt(); err != nil {
		panic(err)
	}
}
