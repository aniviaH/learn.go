package main

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

type Sphere struct {
	Radius float64
}

type Triangle struct {
	base   float64
	height float64
}

// 表面积
func (s Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

// 体积
func (s Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

func (t Triangle) changeBaseByValue(f float64) {
	t.base = f
	return
}
func (t *Triangle) changeBaseByPointer(f float64) {
	t.base = f
	return
}
func (t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

func (m Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return r
}

func TestStructMethod(t *testing.T) {
	m := Movie{
		Name:   "Spiderman",
		Rating: 3.2,
	}
	fmt.Println(m.summary())

	s := Sphere{
		Radius: 2,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())
}

func TestChangeBase(t *testing.T) {
	triangle := Triangle{
		base:   3,
		height: 1,
	}
	fmt.Println(triangle.area())

	triangle.changeBaseByValue(4)
	fmt.Println(triangle.base)

	triangle.changeBaseByPointer(4)
	fmt.Println(triangle.base)
	fmt.Println(triangle.area())
}
