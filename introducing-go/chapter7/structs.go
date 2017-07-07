package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

type circle struct {
	x, y, r float64
}

func ptrCA(c *circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	y := distance(r.x1, r.y1, r.x2, r.y2)
	return l * y
}

type person struct {
	name string
}

func (p *person) talk() {
	fmt.Println("My name is ", p.name)
}

func (a *android) talk() {
	fmt.Println("Model name is ", a.model)
}

type android struct {
	person
	model string
}

type Shapes interface {
	area() float64
}

func totalArea(shapes ...Shapes) float64 {
	var area float64
	fmt.Println("shapes = ", shapes)
	for _, s := range shapes {
		fmt.Println("value of s %p", s)
		area += s.area()
	}
	return area
}

func main() {
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5
	//c := circle{x: 0, y: 1, r: 5}
	c := new(circle)
	c.x = 10
	c.y = 10
	c.r = 20
	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))
	//fmt.Println(ptrCA(c))
	fmt.Println(c.area())
	r := new(Rectangle)
	r.x1 = 10
	r.x2 = 12
	r.y1 = 14
	r.y2 = 16
	fmt.Println(r.area())
	p := new(person)
	p.name = "prem"
	p.talk()
	a := new(android)
	a.name = "mi"
	a.model = "redmi"
	a.person.talk()
	a.talk()
	//a.person.talk()
	fmt.Println(totalArea(r, c))
}
