package main

import "fmt"

func main() {
	arr := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
	fmt.Printf("%d, %d\n", len(arr), cap(arr))

	fmt.Println("-------------------")
	sl1 := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	sl1 = append(sl1, "Update")
	fmt.Println(sl1)
	fmt.Printf("%T\n", sl1)
	fmt.Printf("%d, %d\n", len(sl1), cap(sl1))
	for k, v := range sl1 {
		fmt.Println(k, v)
	}

	fmt.Println("-------------------")
	//pg 57
	sl2 := []int{0, 1, 2, 3, 4}
	sl3 := make([]int, len(sl2)-1)
	//	for i, x := range sl2[1:] {
	//		sl3[len(sl3)-i-1] = x
	//	}
	//sl3 = append(sl3, sl2[1:]...)

	copy(sl3, sl2[:4])
	sl3[0] = 123
	fmt.Println(sl2, sl3)

	fmt.Println("-------------------")
	m := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, x := range m {
		fmt.Println(x)
	}
	fmt.Printf("%#v\n", m)

	fmt.Println("-------------------")
	a := []int{42, 43, 44, 45, 46}
	b := []int{47, 48, 49, 50, 51}
	c := []int{44, 45, 46, 47, 48}
	d := []int{43, 44, 45, 46, 47}
	xi := [][]int{a, b, c, d}
	fmt.Printf("%#v\n", xi)

	for _, x := range xi {
		for _, y := range x {
			fmt.Printf("%d\t", y)
		}
		fmt.Println()
	}

	fmt.Println("-------------------")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	fmt.Println("-------------------")
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	st := make([]string, 0)
	st = append(st, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(len(st), cap(st))
	fmt.Println(st)

	fmt.Println("-------------------")
	n := [][]string{[]string{"James", "Bond", "Shaken, not stirred"}, []string{"Miss", "Moneypenny", "Im 008."}}
	fmt.Println(n)

}
