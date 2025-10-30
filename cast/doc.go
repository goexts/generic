/*
Package cast provides safe, generic alternatives to Go's standard type assertion.
It simplifies type conversions by offering convenient, single-expression functions
that gracefully handle the `value, ok` idiom, avoiding panics on incorrect type
assertions.

# Usage

A common use case is safely dispatching events of different types:

	type UserCreatedEvent struct{ UserID int }
	type OrderPlacedEvent struct{ OrderID string }

	func HandleEvent(event any) {
		if u, ok := cast.As[UserCreatedEvent](event); ok {
			processUserCreated(u)
			return
		}
		if o, ok := cast.As[OrderPlacedEvent](event); ok {
			processOrderPlaced(o)
			return
		}
		// Optionally, handle or log unknown event types.
	}

Another simple example:

	var myVal any = "hello world"

	// Safely cast `myVal` to a string.
	if str, ok := cast.As[string](myVal); ok {
		fmt.Printf("Successfully casted to string: %s\n", str)
	}

	// Attempt to cast to an incorrect type.
	if _, ok := cast.As[int](myVal); !ok {
		fmt.Println("Failed to cast to int, as expected.")
	}

For more details on specific functions, refer to the function documentation.
*/
package cast
