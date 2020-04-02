package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	serverAddr = ":5050"
)

var (
	buffer  = make([]byte, 512)
	message string
)

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Failed to connect to the server\nExit program")
		os.Exit(0)
	}
	defer conn.Close()

	go recieveMessages(conn)

	for {
		message := inputString()
		conn.Write([]byte(message))
	}
}

func inputString() string {
	msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(msg, "\n", "", -1)
}

func recieveMessages(conn net.Conn) {
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		message = string(buffer[:length])
		fmt.Println(message)
	}
}
