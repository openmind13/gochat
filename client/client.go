package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/openmind13/gochat/conn"
	"github.com/openmind13/gochat/message"
	"github.com/openmind13/gochat/view"
)

const (
	serverAddr = ":5050"
)

// Client ...
type Client struct {
	conn     conn.Conn
	username string
}

// NewClient ...
func NewClient(serverAddr string) *Client {
	connection, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil
	}
	c := conn.Conn{connection}
	client := &Client{
		conn: c,
	}
	return client
}

// Start starts client
func (client *Client) Start() error {
	return nil
}

func (client *Client) listenServer() {
	for {
		msg, err := client.conn.ReadMessage()
		if err != nil {
			log.Printf("Error in reading from message\n")
			os.Exit(0)
		}
		switch msg.ControlSequence {
		case message.ControlMessage:
			fmt.Println("have text message")
			view.PrintMessage(msg)
		case message.ControlAccept:
			fmt.Println("have access message")
		case message.ControlRegistration:
			fmt.Println("have registration message")
		}
	}
}

func (client *Client) sender(conn net.Conn) {
	for {
		text := inputString()
		msg := message.Message{
			ControlSequence: message.ControlMessage,
			Author:          client.username,
			Text:            text,
		}
		if err := client.conn.WriteMessage(msg); err != nil {
			fmt.Println("Error in writing message")
			os.Exit(0)
		}
	}
}

func inputString() string {
	msg, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(msg, "\n", "", -1)
}
