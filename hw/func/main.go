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

func main() {
	a := &Circle{Radius: 5}
	a.Scale(25)
	fmt.Println(a.Radius)

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
	WriteString(buffer, "I'm buffer!")
	fmt.Println(buffer.String())
}

func WriteString(w io.Writer, s string) {
	_, err := w.Write([]byte(s))
	if err != nil {
		panic(err)
	}
}
