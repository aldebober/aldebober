package main

import (
	"fmt"
	"math"
)

type square struct {
	l float32
}

type circle struct {
	r float32
}

func (s square) Area() float32 {
	return s.l * s.l
}

func (c circle) Area() float32 {
	return c.r * math.Pi * c.r
}

type shape interface {
	Area() float32
}

func Info(sh shape) {
	fmt.Printf("Area is: %f\n", sh.Area())
}

type person struct {
	first string
	last  string
	age   int
}

type agent struct {
	person
	kill bool
}

func (p person) Speak() {
	fmt.Printf("Name: %v %v, Age: %d\n", p.first, p.last, p.age)
}

func (a agent) Speak() {
	fmt.Printf("Name: %v %v, Age: %d, Allowed to kill: %v\n", a.person.first, a.person.last, a.person.age, a.kill)
}

type people interface {
	Speak()
}

func Tellme(m people) {
	m.Speak()
}

func main() {
	c := circle{13}
	s := square{15}

	Info(c)
	Info(s)

	p := person{
		first: "Yuri",
		last:  "Pavlov",
		age:   45,
	}
	sa := agent{
		person{
			first: "James",
			last:  "Bond",
			age:   49,
		},
		true,
	}

	fmt.Printf("%v\n", p)
	Tellme(p)
	fmt.Printf("%v\n", sa)
	Tellme(sa)
}
