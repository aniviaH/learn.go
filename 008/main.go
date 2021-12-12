package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}
func main() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	//for _, t := range ts {
	//	//defer func(t Test) {
	//	//	t.Close()
	//	//}(t)
	//	t.Close()
	//}

	var t Test
	for i := 0; i < len(ts); i++ {
		t = ts[i]
		//fmt.Println(i)
		//t.Close()       // a b c
		defer t.Close()
	}

	//for _, t := range ts {
	//	//t.Close() // a b c
	//	defer t.Close() // c c c
	//}
}
