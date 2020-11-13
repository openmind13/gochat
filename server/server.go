package main

import (
	"fmt"
	"log"
	"net"

	"github.com/openmind13/gochat/conn"
	"github.com/openmind13/gochat/message"
)

const (
	port = ":5050"

	buffSize = 512
)

// Server ...
type Server struct {
	users         map[User]bool
	listener      net.Listener
	messageBuffer chan message.Message
}

// NewServer create new server
func NewServer(addr string) *Server {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil
	}
	return &Server{
		users:         nil,
		listener:      l,
		messageBuffer: make(chan message.Message),
	}
}

// Start ...
func (s *Server) Start() error {
	go s.sender()
	s.listenSocket()
	return nil
}

// Stop ...
func (s *Server) Stop() {
	s.listener.Close()
}

func (s *Server) listenSocket() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}
		if !s.registerUser(conn) {
			continue
		}
		go s.listenConnection(conn)
	}
}

func (s *Server) listenConnection(connection net.Conn) {
	co := conn.Conn{connection}
	for {
		msg, err := co.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		switch msg.ControlSequence {
		case message.ControlAccept:
		case message.ControlRegistration:
			fmt.Println("have registration message")
		case message.ControlMessage:
			s.messageBuffer <- msg
		}
	}
	co.Close()
}

// sender sends messages from
func (s *Server) sender() {
	for {
		msg, ok := <-s.messageBuffer
		if !ok {
			continue
		}
		s.sendBroadcast(msg)
	}
}

// sendBroadcast sends message to all connected users
func (s *Server) sendBroadcast(msg message.Message) {
	for key := range s.users {
		if err := key.conn.WriteMessage(msg); err != nil {
			s.Stop()
			log.Fatal(err)
		}
	}
}

// registerUser ...
func (s *Server) registerUser(conn net.Conn) bool {
	for {

	}
}

func (s *Server) inChat(user User) bool {
	flag, ok := s.users[user]
	if !ok {
		return false
	}
	return flag
}
