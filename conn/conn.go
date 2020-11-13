package conn

import (
	"errors"
	"net"

	"github.com/openmind13/gochat/message"
)

// Errors
var (
	ErrReadMessage  = errors.New("Error in reading from connection")
	ErrWriteMessage = errors.New("Error in writing message into connection")

	ErrDeserializeMessage = errors.New("Error in deserialing message")
	ErrSerializeMessage   = errors.New("Error in serializing message")
)

// Conn user connection
type Conn struct {
	net.Conn
}

// ReadMessage read from
func (conn *Conn) ReadMessage() (message.Message, error) {
	var buffer []byte
	if _, err := conn.Read(buffer); err != nil {
		return message.Message{}, ErrReadMessage
	}
	msg, err := message.Deserialize(buffer)
	if err != nil {
		return message.Message{}, ErrDeserializeMessage
	}
	return msg, nil
}

// WriteMessage ...
func (conn *Conn) WriteMessage(msg message.Message) error {
	binMsg, err := msg.Serialize()
	if err != nil {
		return ErrSerializeMessage
	}
	if _, err := conn.Write(binMsg); err != nil {
		return ErrWriteMessage
	}
	return nil
}
