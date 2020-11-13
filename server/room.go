package main

// Room ...
type Room struct {
	users []User
}

// UsersCount ...
func (room *Room) UsersCount() int {
	return len(room.users)
}
