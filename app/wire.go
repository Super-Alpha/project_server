// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package app

import (
	"fmt"
	"github.com/google/wire"
)

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return "Hi there!"
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
