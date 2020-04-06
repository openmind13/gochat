package main

import (
	"fmt"
	"log"
	"net"
)

const (
	port = ":5050"
)

// terminal colors
const (
	reset  = "\033[0m"
	red    = "\033[1;31m"
	yellow = "\033[1;33m"
	blue   = "\033[0;34m"
)

var (
	connections = make(map[net.Conn]string)
)

func main() {
	server, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	fmt.Printf("Server started at port: %s", port)

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
	var (
		buffer   = make([]byte, 512)
		nickname string
		message  string
	)

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("error in setting nickname")
		}
		nickname = string(buffer[:length])
		if nicknameIsCorrect(nickname) {
			conn.Write([]byte("ok"))
			connections[conn] = nickname

			log.Printf("| %s%s%s %s%s%s", yellow, nickname, reset, red, "joined the chat", reset)
			//log.Println(nickname + " joined the chat")
			sendToAll(conn, nickname+" joined the chat")

			break
		}
	}

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			log.Printf("| %s%s%s %s%s%s", yellow, nickname, reset, red, "left chat", reset)
			break
		}
		message = string(buffer[:length])

		log.Printf("| %s%s%s > %s", yellow, nickname, reset, message)
		// log.Println(nickname + " > " + message)
		sendToAll(conn, nickname+" > "+message)
	}

	sendToAll(conn, nickname+" left the chat")
	delete(connections, conn)
}

func nicknameIsCorrect(nick string) bool {
	for _, value := range connections {
		if value == nick {
			return false
		}
	}
	return true
}

func sendToAll(sender net.Conn, message string) {
	for conn := range connections {
		if sender != conn {
			conn.Write([]byte(message))
		}
	}
}
