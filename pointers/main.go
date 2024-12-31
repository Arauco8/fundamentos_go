package main

import "fmt"

func main() {
	var i int
	var iP *int

	i = 34
	iP = &i //puntero

	fmt.Println("&i: ", &i)
	fmt.Println("iP: ", iP)
	fmt.Println("*iP: ", *iP)
	fmt.Println("i: ", i)

	*iP = 21

	fmt.Printf("value i: %d, value pointer i: %d\n", i, *iP)
	fmt.Println("-------------------------------------------------------------")
	myVar := 100
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)
	increment(myVar)
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)
	incrementPointer(&myVar)
	fmt.Printf("address: %p, value: %d \n", &myVar, myVar)

}

func increment(val int) {
	val++
	fmt.Println("&val: ", &val)
	fmt.Println("val: ", val)
}

func incrementPointer(val *int) {
	*val++
	fmt.Println("&val: ", &val)
	fmt.Println("val: ", val)
	fmt.Println("*val: ", *val)
}
