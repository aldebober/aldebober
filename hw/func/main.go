package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
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

func tellme(a animals) {
	a.say()
}

func (c cats) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", c.Name, c.Age)
}

func main() {
	c1 := cats{Name: "Motia", Age: 7}
	d1 := dogs{}
	tellme(c1)
	tellme(d1)
	fmt.Println(c1)

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
}

func WriteString(w io.Writer, s string) {
	_, err := w.Write([]byte(s))
	if err != nil {
		panic(err)
	}
}
