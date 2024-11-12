package main

import (
	"fmt"
)

func main() {

	var myIntVar int

	myIntVar = -12

	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint

	myUintVar = 12

	fmt.Println("My variable is: ", myUintVar)

	var myStringVar string

	myStringVar = "Hola ke ace"

	fmt.Println("My variable is: ", myStringVar)

	var myBoolVar bool

	myBoolVar = true

	fmt.Println("My variable is: ", myBoolVar)

	fmt.Println("My variable address is: ", &myStringVar)

	myIntVar2 := 12

	fmt.Println("My variable is: ", myIntVar2)

	myStringVar2 := "My string variable whith :="

	fmt.Println("My variable is: ", myStringVar2)

	myBoolVar2 := true
	fmt.Println("My variable is: ", myBoolVar2)

}
