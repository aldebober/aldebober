package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type user struct {
	First string
	Age   int
}

type agent struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

type Person struct {
	First string
	Age   int
}

type byAge []Person

func (aa byAge) Len() int           { return len(aa) }
func (aa byAge) Swap(i, j int)      { aa[i], aa[j] = aa[j], aa[i] }
func (aa byAge) Less(i, j int) bool { return aa[i].Age < aa[j].Age }

type byName []Person

func (aa byName) Len() int           { return len(aa) }
func (aa byName) Swap(i, j int)      { aa[i], aa[j] = aa[j], aa[i] }
func (aa byName) Less(i, j int) bool { return aa[i].First < aa[j].First }

func main() {

	fmt.Println("----------------")
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	p := []user{u1, u2, u3}
	pb, er := json.Marshal(p)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(string(pb))

	fmt.Println("----------------")
	s1 := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s1)
	bs := []byte(s1)

	var a []agent
	err := json.Unmarshal(bs, &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	for i, person := range a {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}

	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println("----------------")

	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}
	sort.Sort(byAge(people))
	fmt.Println(people)
	sort.Sort(byName(people))
	fmt.Println(people)
}
