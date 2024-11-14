package main

import "fmt"

func main() {

	myArrayVar := [5]int{3, 6, 9, 10, 16}
	fmt.Println("Array: ", myArrayVar, " - len: ", len(myArrayVar))

	mySliceVar := []int{}

	mySliceVar = append(mySliceVar, 12, 34, 54, 31, 12)
	fmt.Println("Slice: ", mySliceVar, " - len: ", len(mySliceVar))

	fmt.Println(myArrayVar[2:4])
	mySliceVar2 := myArrayVar[2:4]
	mySliceVar2[0] = 19                                                                                // tb modifica el valor del array original
	fmt.Println("Slice: ", mySliceVar2, " - len: ", len(mySliceVar2), " - address: ", &mySliceVar2[0]) // toma la referencia del array original por eso son iguales

	fmt.Println("Array: ", myArrayVar, " - len: ", len(myArrayVar), " - address: ", &myArrayVar[2])

	mySliceVar3 := mySliceVar[0:4]
	fmt.Println("Slice: ", mySliceVar3, " - len: ", len(mySliceVar3), " - address: ", &mySliceVar3[0])
	fmt.Println("mySliceVar: ", mySliceVar, " - len: ", len(mySliceVar), " - address: ", &mySliceVar[0])

	mySliceVar4 := make([]int, 3)
	fmt.Println("Slice: ", mySliceVar4, " - len: ", len(mySliceVar4))

	mySliceVar5 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice: ", mySliceVar5, " - len: ", len(mySliceVar5))

}
