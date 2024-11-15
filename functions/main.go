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
	fmt.Println()

	xVal, yVal := function.Split(40)
	fmt.Println(xVal, yVal)
	fmt.Println()

	val2 := function.MSum(1, 2, 3, 1, 2, 3, 4)
	fmt.Println(val2)

}
