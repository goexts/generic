// Package generic demonstrates safe casting dispatch with cast.As (Report Q6).
package examples

import (
	"fmt"

	"github.com/goexts/generic/cast"
)

type UserCreatedEvent struct{ UserID int }
type OrderPlacedEvent struct{ OrderID string }

func HandleEvent(event any) {
	processUserCreated := func(e UserCreatedEvent) { fmt.Println("user:", e.UserID) }
	processOrderPlaced := func(e OrderPlacedEvent) { fmt.Println("order:", e.OrderID) }

	if u, ok := cast.Try[UserCreatedEvent](event); ok {
		processUserCreated(u)
		return
	}
	if o, ok := cast.Try[OrderPlacedEvent](event); ok {
		processOrderPlaced(o)
		return
	}
}

func ExampleHandleEvent() {
	HandleEvent(UserCreatedEvent{UserID: 123})
	HandleEvent(OrderPlacedEvent{OrderID: "A-1"})
	HandleEvent(42)

	// Output:
	// user: 123
	// order: A-1
}
