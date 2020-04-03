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

// terminal colors
const (
	reset  = "\033[0m"
	red    = "\033[1;31m"
	yellow = "\033[1;33m"
	blue   = "\033[0;34m"
)

var (
	buffer   = make([]byte, 512)
	message  string
	nickname string
)

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Failed to connect to the server\nExit program")
		os.Exit(0)
	}

	setNickname(conn)

	go recieve(conn)

	for {
		message = inputString()
		conn.Write([]byte(message))
	}
}

func setNickname(conn net.Conn) {
	var (
		requestNick string
		answer      string
	)
	for {
		fmt.Printf("%sEnter your nickname > %s", red, reset)
		requestNick = inputString()
		conn.Write([]byte(requestNick))

		length, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error in setNickname")
			log.Fatal(err)
		}

		answer = string(buffer[:length])
		if answer == "ok" {
			nickname = requestNick
			return
		}

		fmt.Println("Nickname is used or invalid. Try again")
	}
}

func inputString() string {
	msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(msg, "\n", "", -1)
}

func recieve(conn net.Conn) {
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("%sServer stopped!%s", red, reset)
			os.Exit(0)
		}
		message = string(buffer[:length])
		log.Printf("| %s%s%s", blue, message, reset)
	}
}

// func send(conn net.Conn) {
// 	for {
// 		message = inputString()
// 		conn.Write([]byte(message))
// 	}
// }
