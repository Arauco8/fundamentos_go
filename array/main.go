package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var myIntVar int
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d\n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar1 [5]int
	fmt.Println(myArrayVar1, " - size: ", len(myArrayVar1))

	myArrayVar2 := [3]string{"value1", "value2", "value3"}
	fmt.Println(myArrayVar2, " - size: ", len(myArrayVar2))

	myArrayVar1[0] = 2
	myArrayVar1[1] = 5
	myArrayVar1[2] = 9
	fmt.Println(myArrayVar1, " - size: ", len(myArrayVar1))
	fmt.Println(myArrayVar1[4])
	fmt.Printf("type: %T, value: %v, bytes: %d, bits: %d\n", myArrayVar1, myArrayVar1, unsafe.Sizeof(myArrayVar1), unsafe.Sizeof(myArrayVar1)*8)

	for i := range myArrayVar1 {
		fmt.Println("index: ", i, " - value: ", myArrayVar1[i])
	}

	for i, v := range myArrayVar1 { // devuelve el indice en i y el valor en v
		fmt.Println("index: ", i, " - value: ", v)
	}

}
