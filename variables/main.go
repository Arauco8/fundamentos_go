package main

import (
	"fmt"
	"strconv"
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

	/*

		int8     8 bits  -128 to 127
		int16    16 bits -2^15 to 2^15 -1
		int32    32 bits -2^31 to 2^31 -1
		int 64   64 bits -2^63 to 2^63 -1
		int Platform dependent Platform dependent

	*/

	var my8BitsUnitVar uint8 = 20
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my8BitsUnitVar, my8BitsUnitVar, unsafe.Sizeof(my8BitsUnitVar), unsafe.Sizeof(my8BitsUnitVar)*8)

	var my16BitsUnitVar uint16 = 30
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my16BitsUnitVar, my16BitsUnitVar, unsafe.Sizeof(my16BitsUnitVar), unsafe.Sizeof(my16BitsUnitVar)*8)

	var my32BitsUnitVar uint32 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my32BitsUnitVar, my32BitsUnitVar, unsafe.Sizeof(my32BitsUnitVar), unsafe.Sizeof(my32BitsUnitVar)*8)

	var my64BitsUnitVar uint64 = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", my64BitsUnitVar, my64BitsUnitVar, unsafe.Sizeof(my64BitsUnitVar), unsafe.Sizeof(my64BitsUnitVar)*8)

	var mUinitVar uint = 90
	fmt.Printf("type: %T, value: %d, bytes: %d, bits: %d \n", mUinitVar, mUinitVar, unsafe.Sizeof(mUinitVar), unsafe.Sizeof(mUinitVar)*8)

	var myFloat32Var float32 = 3.14
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64 = 590.12435
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myFloat64Var, myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	myOtherFloat := 7654.123
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d \n", myOtherFloat, myOtherFloat, unsafe.Sizeof(myOtherFloat), unsafe.Sizeof(myOtherFloat)*8)

	var myStringVar3 string = "test"           // string son secuencias de bytes
	fmt.Printf("mi valor %s \n", myStringVar3) //%s definir que estamos recibiendo un valor de tipo string

	myStringVar4 := `Mi variable de tipo texto en Go
	con multiples
	lineas
	:)`

	fmt.Printf("mi valor es: %s \n", myStringVar4)

	{

		floatVar := 33.11
		fmt.Printf("type: %T, value: %f\n", floatVar, floatVar)

		floatStrVar := fmt.Sprintf("%.4f", floatVar) //en vez de imprimir le asigna eso a la variable, y cuando usamos %f agregando .4 le damos formato
		fmt.Printf("type: %T, value: %s\n", floatStrVar, floatStrVar)

		intVar := 37
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)

		intStrVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStrVar, intStrVar)

		intVar2 := 342
		intStrVar2 := strconv.Itoa(intVar2)
		fmt.Printf("type: %T, value: %s\n", intStrVar2, intStrVar2)

		strIntVar2, err := strconv.Atoi("15")
		fmt.Printf("type: %T, value: %d, err: %v\n", strIntVar2, strIntVar2, err)

		strIntVar3, _ := strconv.ParseInt("10", 10, 64)
		fmt.Printf("type: %T, value: %d\n", strIntVar3, strIntVar3) // agrego _ para omitir usar err

		strFloatVar, err := strconv.ParseFloat("-11.2", 64)
		fmt.Printf("type: %T, value: %f, err: %v\n", strFloatVar, strFloatVar, err)
	}

}
