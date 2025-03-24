package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func main() {
	a := &Circle{Radius: 5}
	a.Scale(25)
	fmt.Println(a.Radius)
}
