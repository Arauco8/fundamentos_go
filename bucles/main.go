package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i < 10; i++ {

		sum++

	}

	fmt.Println(sum)

	sum = 1

	for sum < 1000 {
		sum++
	}

	fmt.Println(sum)

	for {
		if sum > 1000 {
			break
		}
		sum++
	}

	fmt.Println(sum)
}
