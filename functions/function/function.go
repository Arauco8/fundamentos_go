package function

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Display(myValue int) {
	fmt.Println()
	fmt.Printf("type: %T, value: %d, address: %v\n", myValue, myValue, &myValue)
	fmt.Println()
}

func Add(x, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value)
	}
	fmt.Println()
}

func Calc(op Operation, x, y float64) (float64, error) {
	switch op {
	case SUM:
		return x + y, nil
	case SUB:
		return x - y, nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y mustn't be zero")
		}
		return x / y, nil
	case MUL:
		return x * y, nil
	}

	return 0, errors.New("has been an error")
}

func Split(v int) (x, y int) { //como ya tenemos nombrado los valores que va retornar no hace falta agregarlos en el return
	x = v * 4 / 9
	y = v - x
	return
}

func MSum(values ...float64) float64 { //... con esto internamente lo manejamos como un array, elipsis
	var sum float64
	for _, v := range values {
		sum += v
	}

	return sum

}
