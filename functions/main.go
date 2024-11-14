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
}
