package main

import (
	"fmt"
	"reflect"
)

type Movie struct {
	Name         string
	Rating       float64
	privateMoney int64
}

type Address struct {
	Number int
	Street string
	City   string
}

type SuperHero struct {
	Name    string
	Age     int
	Address Address
}

type Alarm struct {
	Time  string
	Sound string
}

func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Kanon",
	}
	return a
}

type int1 int
type int2 = int

func main() {
	{
		m := Movie{
			Name:   "哈利波特",
			Rating: 1999.0,
		}
		//m := new(Movie)
		fmt.Println(m, m.Name, m.Rating)

		// %v 与 %+v
		var m2 Movie
		fmt.Printf("%+v\n", m2)
		fmt.Printf("%v\n", m2)

		fmt.Printf("%+v\n", m)
		fmt.Printf("%v\n", m)
	}

	{
		e := SuperHero{
			Name: "Batman",
			Age:  32,
			Address: Address{
				Number: 1122,
				Street: "xxx街道",
				City:   "xxx城市",
			},
		}
		fmt.Printf("%+v\n", e)
		fmt.Println(e.Address.Street)
		fmt.Println(*(&e.Address))
	}

	{
		fmt.Printf("%+v\n", NewAlarm("7:30"))
	}

	{
		m1 := Movie{
			Name:   "哈利波特",
			Rating: 1,
		}
		m2 := Movie{
			Name:   "哈利波特",
			Rating: 2,
		}
		a1 := Alarm{
			Time:  "7:30",
			Sound: "1",
		}
		if m1 == m2 {
			fmt.Println("m1 and m2 are the same")
		} else {
			fmt.Println("m1 and m2 not same")
		}

		//if m1 == m3 { // 不同类型的结构体进行比较，会编译报错
		//	fmt.Println("m1 and a1 are the same")
		//} else {
		//	//
		//}
		// 试图比较两个结构体之前，必须确定它们的类型相同
		m1Type := reflect.TypeOf(m1)
		//m2Type := reflect.TypeOf(m2)
		a1Type := reflect.TypeOf(a1)
		if m1Type == a1Type {
			fmt.Println("m1 and a1 are the same type")
		} else {
			fmt.Println("m1 and a1 are not same type")
		}

		fmt.Printf("%+v\n", m1)
		fmt.Printf("%+v\n", m2)
		fmt.Println("m1 type:", m1Type)
		fmt.Println("a1 type:", a1Type)
	}

	{
		m3 := Movie{}
		m3.Name = "复仇者联盟"
		fmt.Println(m3, &m3, *&m3.Name)

		m4 := m3
		fmt.Println("m4 == m3", m3 == m4)
		m4.Name = "蜘蛛侠"
		//fmt.Println(m3, &m3)
		//fmt.Println(m4, &m4)
		fmt.Printf("%+v\n", m3)
		fmt.Printf("%+v\n", m4)
		fmt.Printf("%p\n", &m3)
		fmt.Printf("%p\n", &m4)

		p := &m3
		p.Name = "钢铁侠"
		//fmt.Printf("%+v\n", m3)
		//fmt.Printf("%p\n", &m3)
		//fmt.Println(p.Name)
		//fmt.Println(*p == m3)
		//fmt.Println(*p == *&m3)
		fmt.Printf("%+v\n", m3)
		fmt.Printf("%+v\n", *p)
		fmt.Printf("%p\n", &m3)
		fmt.Printf("%p\n", p)
		fmt.Println(p.Name, *p)
	}

	{
		var a int1
		var b int2
		var c int
		fmt.Println(a, b, c)
	}

}
