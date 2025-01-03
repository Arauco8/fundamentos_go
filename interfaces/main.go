package main

import (
	"fmt"
	"proyectos/fundamentos_go/interfaces/vehicles"
)

//Con las Interfaces podemos definir un contrato que deben cumplir los tipos que la implementan. Implementar Polimorfismo
//Las interfaces son un tipo de dato abstracto, no se pueden instanciar

func main() {

	Display("Hola Mundo")
	Display(10)
	Display(10.5)
	Display(true)
	fmt.Println("----------------------------------------------------")

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "MOTORCYCLE", "TRUCK", "CAR"}

	var d float64
	for _, v := range vArray {
		fmt.Printf("Vehicle: %s\n", v)
		vehicle, err := vehicles.NewVehicle(v, 400)
		if err != nil {
			fmt.Println(err)
			fmt.Println()
			continue //vuelve al inicio del ciclo y omite el resto del codigo
		}
		distance := vehicle.Distance()
		fmt.Printf("Vehicle: %s, Distance: %.2f\n", v, distance)
		fmt.Println()
		d += distance
	}

	fmt.Printf("Total Distance: %.2f\n", d)
	fmt.Println()
}

func Display(value interface{}) { //Recibe cualquier tipo de dato
	fmt.Println(value)
}
