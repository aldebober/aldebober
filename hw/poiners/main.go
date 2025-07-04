package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walk.")
}

func (d *dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

type first struct {
	name string
}

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)

	d1 := &dog{"Henry"}
	youngRun(d1)

	d2 := &dog{"Padget"}
	youngRun(d2)
	d2 = d2.changeName("Rover")
	youngRun(d2)

	n1 := first{"yurix"}
	name_v(n1)
	fmt.Println(n1.name)
	name_p(&n1)
	fmt.Println(n1.name)

}

func name_v(s first) {
	s.name = "pavlov"
}

func name_p(s *first) {
	s.name = "bob"
}

func (d *dog) changeName(s string) *dog {
	d.first = s
	return d
}
