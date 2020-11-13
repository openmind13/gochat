package main

import (
	"errors"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server := NewServer(":5000")
	if server == nil {
		return errors.New("Server not created")
	}
	if err := server.Start(); err != nil {
		return err
	}
	return nil
}
