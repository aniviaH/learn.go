package main

import "fmt"

type Movie struct {
	Name   string
	Rating float64
}

func main() {
	m := Movie{
		Name:   "哈利波特",
		Rating: 1999.0,
	}
	fmt.Println(m, m.Name, m.Rating)

	// %v 与 %+v
	var m2 Movie
	fmt.Printf("%+v\n", m2)
	fmt.Printf("%v\n", m2)

	fmt.Printf("%+v\n", m)
	fmt.Printf("%v\n", m)
}
