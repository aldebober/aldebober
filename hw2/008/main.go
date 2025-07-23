package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

type Person struct {
	Name string
	Age  int16
}

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"fc": DoubleAge,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.func"))
}

func DoubleAge(a int16) int16 {
	return a * 2
}

func main() {
	p1 := Person{"Lenin", 79}
	p2 := Person{"Hitler", 63}

	people := []Person{p1, p2}
	data := struct {
		P []Person
	}{
		people,
	}

	fmt.Println(data)
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.func", data)
	if err != nil {
		log.Fatal(err)
	}
}
