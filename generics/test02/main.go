package main

import (
	"fmt"
)

func SumInts(m map[string]int64) int64 {
	var res int64
	for _, item := range m {
		res += item
	}

	return res
}

func SumFloat64(m map[string]float64) float64 {
	var res float64
	for _, item := range m {
		res += item
	}

	return res
}

func SumGenerics[K comparable, V int64 | float64](m map[K]V) V {
	var res V
	for _, item := range m {
		res += item
	}

	return res
}

type Number interface {
	int64 | float64
}

func SumGenericsV2[K comparable, V Number](m map[K]V) V {
	var res V
	for _, item := range m {
		res += item
	}

	return res
}

func main() {

	intMap := map[string]int64{
		"a": 1,
		"b": 2,
	}

	floatMap := map[string]float64{
		"a": 1.1,
		"b": 1.2,
	}

	fmt.Printf("summary of int: %d, summary of float: %f\n",
		SumInts(intMap), SumFloat64(floatMap))

	fmt.Printf("generics summary of int: %d, sumamry of float: %f\n",
		SumGenerics[string, int64](intMap), SumGenerics[string, float64](floatMap))

	fmt.Printf("generics summary of int: %d, summary of float: %f\n",
		SumGenerics(intMap), SumGenerics(floatMap))

	fmt.Printf("generics v2 summary of int: %d, summary of float: %f\n",
		SumGenericsV2(intMap), SumGenericsV2(floatMap))

}
