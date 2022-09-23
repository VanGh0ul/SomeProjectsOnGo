package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 19

	fmt.Println("first map:", m)

	v1 := m["k1"]
	fmt.Println("val1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("new map after delete:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2, "clash": 3}
	fmt.Println("init new map:", n)
}
