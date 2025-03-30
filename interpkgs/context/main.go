package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background() //return context
	ctx = context.WithValue(ctx, "key", "value")

	// Use the context in a function
	value := getValueFromContext(ctx, "key")
	fmt.Println("Value from context:", value)

	// timeout example
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Create a context with a timeout 5segundos
	defer cancel()                                                           // Cancel the context when done, se ejecuta al final de la función

	go myProcess(ctx2) // Start a goroutine that uses the context
	// Simulate some work in the main goroutine

	select {
	case <-ctx2.Done(): // Wait for the context to be done, like the channel
		fmt.Println("Context done:", ctx2.Err())
		/*case <-time.After(3 * time.Second): // Simulate some work
			fmt.Println("Work done before timeout")
		case <-time.After(6 * time.Second): // Simulate some work that takes too long
			fmt.Println("Work done after timeout")*/
	}
}

func getValueFromContext(ctx context.Context, key string) string {
	// Retrieve the value from the context
	value := ctx.Value(key)
	if value == nil {
		return "No value found"
	}
	return value.(string)
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ohh, time out!")
			return
		default:
			// Simulate some work
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond) // Sleep for 500 milliseconds
	}
}
