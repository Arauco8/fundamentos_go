package main

import "fmt"

func main() {

	yearsOld := 16

	if yearsOld > 18 {
		fmt.Printf("%d es mayor a 18\n", yearsOld)
	}

	boolVar := true

	if boolVar {
		fmt.Println("es Verdadero")
	} else {
		fmt.Println("es false")
	}

	if value := true; value {
		fmt.Println("es verdadero")
	}

	number := 3

	if number == 1 {
		fmt.Println("uno")
	} else if number == 2 {
		fmt.Println("dos")
	} else if number == 3 {
		fmt.Println("tres")
	} else {
		fmt.Println("ninguna es valida")
	}
}
