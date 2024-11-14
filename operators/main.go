package main

import "fmt"

func main() {

	yearsOld := 33

	fmt.Println(yearsOld > 30)
	fmt.Println(yearsOld < 33)
	fmt.Println(yearsOld <= 33)
	fmt.Println(yearsOld >= 40)
	fmt.Println(yearsOld == 33)

	fmt.Println("---- OR ----")
	fmt.Println(yearsOld < 33 || yearsOld == 33)
	fmt.Println(yearsOld < 33 || yearsOld == 34)
	fmt.Println(yearsOld < 40 || yearsOld == 33)
	fmt.Println("---- AND ----")
	fmt.Println(yearsOld < 33 && yearsOld == 33)
	fmt.Println(yearsOld < 33 && yearsOld == 34)
	fmt.Println(yearsOld < 40 && yearsOld == 33)
	fmt.Println("---- NOT ----")
	fmt.Println(true)
	fmt.Println(!true)
	fmt.Println(yearsOld < 40)
	fmt.Println(!(yearsOld < 40))
	fmt.Println("---- OPERATOR COMPUESTO ----")
	fmt.Println(yearsOld < 25 && yearsOld == 33 || yearsOld < 40)   //true
	fmt.Println(yearsOld < 25 && (yearsOld == 33 || yearsOld < 40)) //false
}
