package main

import (
	"fmt"
	"unsafe"
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

	const myIntConst int = 12
	fmt.Println("Mi constante es: ", myIntConst)

	const myFirstStringConst = "a12" // la constante van a ser solo de lectura
	fmt.Println("mi constante es: ", myFirstStringConst)

	const myStringConst string = "test"
	fmt.Println("Mi constante es: ", myStringConst)

	myOtherScopeVariable := 50

	{

		myScopeVariable := 40 //solamente vive dentro de este scope {}
		{
			fmt.Println("Mi variable: ", myOtherScopeVariable)
			fmt.Println("Mi variable: ", myScopeVariable)
		}

	}

	fmt.Println("Mi variable: ", myOtherScopeVariable)

	var my8BitsUnitVar uint8 = 20
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my8BitsUnitVar, my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

	var my16BitsUnitVar uint16 = 30
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my16BitsUnitVar, my16BitsUnitVar, unsafe.Sizeof(my16BitsUnitVar), unsafe.Sizeof(my16BitsUnitVar)*8)

	var my32BitsUnitVar uint32 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my32BitsUnitVar, my32BitsUnitVar, unsafe.Sizeof(my32BitsUnitVar), unsafe.Sizeof(my32BitsUnitVar)*8)

	var my64BitsUnitVar uint64 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my64BitsUnitVar, my64BitsUnitVar, unsafe.Sizeof(my64BitsUnitVar), unsafe.Sizeof(my64BitsUnitVar)*8)

}
