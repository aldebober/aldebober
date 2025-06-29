package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

type Circle struct {
	Radius float64
}

func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius
}

type dogs struct{}

type cats struct {
	Name string
	Age  int
}

func (d dogs) say() {
	fmt.Println("Wow")
}

func (c cats) say() {
	fmt.Println("Mew")
}

type animals interface {
	say()
}

type Shape interface {
	Area() float32
}

func tellme(a animals) {
	a.say()
}

func Info(h Shape) {
	fmt.Println(h.Area())
}

func (c cats) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", c.Name, c.Age)
}

type Square struct {
	length float32
	width  float32
}

type Cycle struct {
	radius float32
}

func (s Square) Area() float32 {
	return s.length * s.width
}

func (c Cycle) Area() float32 {
	return math.Pi * c.radius * 2
}

func main() {
	square1 := Square{12, 51}
	cycle1 := Cycle{3.5}

	Info(square1)
	Info(cycle1)

	c1 := cats{Name: "Motia", Age: 7}
	d1 := dogs{}
	tellme(c1)
	tellme(d1)
	fmt.Println(c1)
	defer fmt.Println("The END")

	a := &Circle{Radius: 5}
	a.Scale(25)
	fmt.Println(a.Radius)
	fmt.Println(a.Area())

	file, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []byte("Hello, World!")
	fmt.Println(string(data))
	//data[0] = 15
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(data)
	buffer.WriteString("Golang here")
	fmt.Println(buffer.String())
	buffer.Reset()

	WriteString(file, "output to file")
	WriteString(file, "\n--------------\n")
	WriteString(buffer, "I'm buffer!")
	fmt.Println(buffer.String())

	str := "World!"
	byte := []byte(str)

	fmt.Println(byte)
	byte[2] = 40
	fmt.Println(string(byte))

	foodata := []int{12, 32, 14, 11, 91}
	fmt.Println(foo(23, 56, 78))
	fmt.Println(bar(foodata))

	f1 := func(s string) { fmt.Printf("My name is %s\n", s) }
	f1("Yura")
	f2 := caller()
	f2(42)

	fmt.Println(process(add, 3, 6))
	TimedFunction(bar, foodata)

	CallbackPrintnum(foodata, printnum)
}

func CallbackPrintnum(data []int, f func(n int)) {
	for _, i := range data {
		f(i)
	}
}

func printnum(n int) {
	fmt.Println(n)
}

func caller() func(i int) {
	a := func(i int) { fmt.Printf("%d\n", i) }
	return a
}

type Callback func(i int, x int) int

func process(cb Callback, a int, b int) int {
	return cb(a, b)

}

func TimedFunction(fn func(i []int) int, data []int) {
	start := time.Now()
	fn(data)
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)
}

func WriteString(w io.Writer, s string) {
	_, err := w.Write([]byte(s))
	if err != nil {
		panic(err)
	}
}

func foo(i ...int) int {
	sum := 0
	for _, x := range i {
		sum += x
	}

	return sum
}

func bar(i []int) int {
	sum := 0
	for _, i := range i {
		sum += i
	}

	return sum
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}
