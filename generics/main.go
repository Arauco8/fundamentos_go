package main

import (
	"fmt"

	"cmp"
)

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
	anyTypeTwo(1, "5")
	comparableType(1, 1)
	orderedValues(2, 4)
	fmt.Println()
	csInt := CustomSlice[int]{1, 2, 3, 4}
	fmt.Println(csInt)
	csString := CustomSlice[string]{"a", "b", "c", "4"}
	fmt.Println(csString)
	fmt.Println()
	x, y := Point(5), Point(2)
	fmt.Println(MinNumber(x, y))
	fmt.Println()
	df := MyGenericStruct[myFirstData]{StrValue: "First", Data: myFirstData{}}
	df.Data.PrintOne()
	ds := MyGenericStruct[mySecondData]{StrValue: "Second", Data: mySecondData{}}
	ds.Data.PrintTwo()
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

func anyTypeTwo[T1 any, T2 any](v T1, v2 T2) {
	fmt.Println(v, v2)
}

func comparableType[T comparable](v, v2 T) {
	fmt.Println(v, v2)
	fmt.Println(v != v2) // nose puede comparar con type any
}

func orderedValues[T cmp.Ordered](v, v2 T) {
	fmt.Println(v, v2)
	fmt.Println(v != v2)
	fmt.Println(v < v2)
	fmt.Println(v > v2)
}

type CustomSlice[T int | string] []T

func MinNumber[T N1](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}

type Point int

type N1 interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 // ~ means "any type except"
}

type (
	MyGenericStruct[T any] struct {
		StrValue string
		Data     T
	}
	myFirstData  struct{}
	mySecondData struct{}
)

func (d myFirstData) PrintOne() {
	fmt.Println("PrintOne")
}

func (d mySecondData) PrintTwo() {
	fmt.Println("PrintTwo")
}
