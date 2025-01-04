package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidStr := uuid.New().String()
	fmt.Println(uuidStr)

	id := uuid.New()
	fmt.Println(id.Version())
	fmt.Println(id.Variant())
}
