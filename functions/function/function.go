package function

import (
	"fmt"
)

func Display(myValue int) {
	fmt.Println()
	fmt.Printf("type: %T, value: %d, address: %v\n", myValue, myValue, &myValue)
	fmt.Println()
}

func Add(x, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value)
	}
	fmt.Println()
}
