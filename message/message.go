package message

import (
	"encoding/json"
)

// Control sequences for messages
var (
	ControlMessage      = uint8(0)
	ControlAccept       = uint8(1)
	ControlRegistration = uint8(2)
)

// Message communication struct
type Message struct {
	ControlSequence uint8  `json:"controlsequence"`
	Author          string `json:"user"`
	Room            string `json:"room"`
	Text            string `json:"text"`
}

// Serialize convert Message to byte sequences
func (message *Message) Serialize() ([]byte, error) {
	buffer, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}

// Deserialize convert bytes to Message
func Deserialize(buffer []byte) (Message, error) {
	var msg Message
	if err := json.Unmarshal(buffer, &msg); err != nil {
		return Message{}, err
	}
	return msg, nil
}
