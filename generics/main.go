package main

import "fmt"

func main() {
	v1 := []float64{1.3, 5.45, 12.223, 6.92, 78.102}
	v2 := []int32{9, 23, 1, 23, 8, 98}

	println(sum01(v1))
	println(sum01(v2))
	fmt.Println()
	fmt.Println(sum02(v1))
	fmt.Println(sum02(v2))
	fmt.Println()
	anyType(1, 1)
	anyType("1", "3")
}

func sum01[N int | int32 | int64 | float32 | float64](v []N) N {
	var sum N
	for _, val := range v {
		sum += val
	}
	return sum
}

type Number interface {
	int | int32 | int64 | float32 | float64
}

func sum02[N Number](v []N) N {
	var sum N
	for _, val := range v {
		sum += val
	}
	return sum
}

func anyType[T any](v, v2 T) {
	fmt.Println(v, v2)
}
