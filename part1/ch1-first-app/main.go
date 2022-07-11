package main

import (
	"partyinvites/handler"
)

type Rsvp struct {
	Name       string
	Email      string
	Phone      string
	WillAttend bool
}

var responses = make([]*Rsvp, 0, 10)

func main() {
	handler.LoadTemplates()
	handler.Handler()
}
