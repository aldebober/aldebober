package main

import "fmt"

type person struct {
	first       string
	last        string
	ice_creames []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		ice_creames: []string{
			"chokolad",
			"banana",
		},
	}

	p2 := person{
		first: "Monney",
		last:  "Penny",
		ice_creames: []string{
			"vanila",
			"strawberry",
		},
	}

	for _, f := range p1.ice_creames {
		fmt.Println(p1.first, p1.last, f)
	}
	for _, f := range p2.ice_creames {
		fmt.Println(p2.first, p2.last, f)
	}
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)

	fmt.Println("-----------------")
	pp := map[string]person{
		p1.first: p1,
		p2.first: p2,
	}
	for n, p := range pp {
		fmt.Printf("%s\t%#v\n", n, p)
	}

	fmt.Println("-----------------")
	type engine struct {
		electric bool
	}
	type vehicle struct {
		engine engine
		make   string
		model  string
		doors  int
		color  string
	}

	v1 := vehicle{
		engine: engine{electric: true},
		make:   "USA",
		model:  "Tesla",
		doors:  5,
		color:  "red",
	}
	v2 := vehicle{
		engine: engine{electric: false},
		make:   "EU",
		model:  "WV",
		doors:  3,
		color:  "blue",
	}

	fmt.Println(v1.model, v1.engine)
	fmt.Println(v2.model, v2.engine)

	fmt.Println("-----------------")
	a := struct {
		first     string
		friend    map[string]int
		favDrinks []string
	}{
		first: "yuri",
		friend: map[string]int{
			"cats":  2,
			"human": 1,
		},
		favDrinks: []string{"beer", "vine"},
	}

	fmt.Printf("%s\t%#v\n", a.first, a)
}
