// Package main demonstrates safe casting dispatch with cast.As (Report Q6).
package main

import (
	"fmt"

	"github.com/goexts/generic/cast"
)

type UserCreatedEvent struct{ UserID int }
type OrderPlacedEvent struct{ OrderID string }

func processUserCreated(e UserCreatedEvent) { fmt.Println("user:", e.UserID) }
func processOrderPlaced(e OrderPlacedEvent) { fmt.Println("order:", e.OrderID) }

func HandleEvent(event any) { // Q6
	if u, ok := cast.As[UserCreatedEvent](event); ok {
		processUserCreated(u)
		return
	}
	if o, ok := cast.As[OrderPlacedEvent](event); ok {
		processOrderPlaced(o)
		return
	}
	// ignore others
}

func main() {
	HandleEvent(UserCreatedEvent{UserID: 123})
	HandleEvent(OrderPlacedEvent{OrderID: "A-1"})
	HandleEvent(42)
}
