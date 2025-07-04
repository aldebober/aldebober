package main

import "fmt"

type Nums interface {
	int | float64
}

func summ[K comparable, T Nums](m map[K]T) T {
	var s T
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	a := map[string]int{
		"first": 34,
		"last":  54,
	}
	b := map[string]float64{
		"first": 34.52,
		"last":  87.53,
	}

	fmt.Println(summ(a))
	fmt.Println(summ(b))
}
