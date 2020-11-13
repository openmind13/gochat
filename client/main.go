package main

import (
	"log"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	time.Sleep(10 * time.Second)

	return nil
}
