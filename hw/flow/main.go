package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {

	x := rand.Intn(250)
	fmt.Println("var x is: ", x)

	switch {
	case x <= 100:
		fmt.Println("x in 0 - 100")
	case x > 100 && x <= 200:
		fmt.Println("x in 101 - 200")
	default:
		fmt.Println("x in 201 - 250")
	}

	fmt.Println("-------------------")
	x = rand.Intn(10)
	y := rand.Intn(10)

	fmt.Println("var x is: ", x)
	fmt.Println("var y is: ", y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y less 4")
	case x > 6 && y > 6:
		fmt.Println("x and y bigger 6")
	case x >= 4 && x <= 6:
		fmt.Println("x in 4 - 6")
	case y != 5:
		fmt.Println("y is not 5")
	default:
		fmt.Println("None")
	}

	fmt.Println("-------------------")

	for i := 0; i <= 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("x is %d\t %d\n", x, i)
	}

	fmt.Println("-------------------")
	for x := 0; x < 21; x++ {
		if x%2 != 0 {

			continue
		}
		fmt.Println(x)
	}

	fmt.Println("-------------------")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d:\t", i)
		for y := 0; y < 5; y++ {
			fmt.Printf("%d\t", y)
		}
		fmt.Println()
	}

	fmt.Println("-------------------")
	xi := []string{"😃", "🌴", "🌞", "🔆", "🌻", "🙌", "🌴", "🌴", "🌴", "🌴", "🌴", "🌴", "🌴", "🌴"}
	for i, x := range xi {
		fmt.Printf("%d:\t%s (%#v)\n", i, x, x)
	}

	fmt.Println("-------------------")
	m := map[string]int{
		"tony": 42,
		"max":  31,
	}
	for i, x := range m {
		fmt.Printf("%s, %d\n", i, x)
	}
	fmt.Println(m["tony"])
}
