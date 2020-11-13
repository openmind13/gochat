package main

import "github.com/openmind13/gochat/conn"

// User ...
type User struct {
	conn conn.Conn
	name string
	auth bool
}
