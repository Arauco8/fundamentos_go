package main

import (
	"fmt"
	"proyectos/fundamentos_go/functions/function"
)

func main() {
	var myIntVar int64
	function.Display(int(myIntVar))

	fmt.Println()
	v := function.Add(4, 2)
	fmt.Println(v)
	fmt.Println()
	function.RepeatString(10, "LA")

	fmt.Println()
	value, err := function.Calc(function.SUM, 20.12, 34)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Println("value: ", value, " - err: ", err)
}
