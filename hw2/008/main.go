package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.func"))
}

var fm = template.FuncMap{
	"FMDate": dayMonthYear,
}

func dayMonthYear(t time.Time) string {
	return t.Format("01-02-2006")
}

func main() {
	t := time.Now()
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.func", t)
	if err != nil {
		log.Fatal(err)
	}
}
