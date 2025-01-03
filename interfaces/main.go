package main

import "fmt"

//Con las Interfaces podemos definir un contrato que deben cumplir los tipos que la implementan. Implementar Polimorfismo
//Las interfaces son un tipo de dato abstracto, no se pueden instanciar

func main() {

	Display("Hola Mundo")
	Display(10)
	Display(10.5)
	Display(true)
}

func Display(value interface{}) { //Recibe cualquier tipo de dato
	fmt.Println(value)
}
