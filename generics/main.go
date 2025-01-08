package main

func main() {
	v1 := []float64{1.3, 5.45, 12.223, 6.92, 78.102}
	v2 := []int32{9, 23, 1, 23, 8, 98}

	println(sum01(v1))
	println(sum01(v2))
}

func sum01[N int | int32 | int64 | float32 | float64](v []N) N {
	var sum N
	for _, val := range v {
		sum += val
	}
	return sum
}
