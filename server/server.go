package main

import (
	"fmt"
	"log"
	"net"
)

const (
	port = ":5050"
)

var (
	connections = make(map[net.Conn]bool)
	buffer      = make([]byte, 512)
	message     string
)

func main() {
	server, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
	// for {
	// 	length, err := conn.Read(buffer)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		break
	// 	}
	// 	message += string(buffer[:length])
	// }
}

func handleConnection(conn net.Conn) {
	connections[conn] = true
	conn.Write([]byte("Welcome to chat server"))

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("client left chat")
			break
		}
		message = string(buffer[:length])
		fmt.Println(message)
		sendToAll(conn, message)
	}
	delete(connections, conn)
}

func sendToAll(sender net.Conn, message string) {
	for conn := range connections {
		if sender != conn {
			conn.Write([]byte(message))
		}
	}
}
