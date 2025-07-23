package main

import (
	"log"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func doubleInt(d int) int {
	return d * 2
}

func main() {
	sl := []string{"one", "two", "three"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gotext", sl)

	if err != nil {
		log.Fatal(err)
	}
	mp := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl.gotext", mp)
	if err != nil {
		log.Fatal(err)
	}

	p1 := Person{
		Name: "Jesus",
		Age:  33,
	}
	p2 := Person{"Budha", 49}
	p3 := Person{"Hitler", 59}

	persons := []Person{p1, p2, p3}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl.struct", persons)
	if err != nil {
		log.Fatal(err)
	}

}
