package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("address: %p \n", &m)
	m.Name = name
}

func (m *MyStruct) SetP(name string) {
	fmt.Printf("address: %p \n", m) //como es un puntero imprimo el puntero donde muestra la direccion de memoria
	m.Name = name
}

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
	fmt.Println("-------------------------------------------------------------")

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)
	fmt.Printf("address: %p, value: %v \n", mySlice, mySlice)
	fmt.Printf("address1: %p, address2: %p, address3: %p \n", &mySlice[0], &mySlice[1], &mySlice[2])
	updateSlice(mySlice)
	fmt.Println(mySlice)

	fmt.Println("-------------------------------------------------------------")
	myStruct := &MyStruct{ID: 1, Name: "Caupolican"} //al hacer & estamos creando un puntero de MyStruct
	fmt.Printf("address: %p \n", &myStruct)
	fmt.Printf("id: %d, name: %s \n", myStruct.ID, myStruct.Name)
	updateStruct(myStruct)
	fmt.Printf("id: %d, name: %s \n", myStruct.ID, myStruct.Name)
	fmt.Println()
	myStruct.Set("Ernesto")
	fmt.Printf("id: %d, name: %s \n", myStruct.ID, myStruct.Name)
	myStruct.SetP("Mario Roberto")
	fmt.Printf("id: %d, name: %s \n", myStruct.ID, myStruct.Name)
}

func increment(val int) {
	val++
	fmt.Println("&val: ", &val)
	fmt.Println("val: ", val)
}

func incrementPointer(val *int) {
	*val++
	fmt.Println("&val: ", &val) //dirección de memoria del puntero, cuando generamos un puntero en go, el puntero es un valor que se guarda en una dirección de memoria
	fmt.Println("val: ", val)   //direccion de memoria de la variable que le estamos pasando, la referencia de la variable
	fmt.Println("*val: ", *val) //valor de la variable que le estamos pasando
}

func updateSlice(slice []int) {
	fmt.Printf("address: %p, value: %v \n", slice, slice)
	fmt.Printf("address1: %p, address2: %p, address3: %p \n", &slice[0], &slice[1], &slice[2])
	slice[0] = 12 //los slices son punteros a un array, por lo que si modificamos un valor de un slice, se modifica el valor del array, trabaja por referencia tambien.
}

func updateStruct(myStruct *MyStruct) { //como hicimos un puntero de MyStruct, estamos pasando la referencia de la variable.
	fmt.Printf("address: %p \n", myStruct)
	myStruct.Name = "Lautaro"
	myStruct.ID = 999
}
