/* Author: Prem karat
   Licnese: MIT
   Adding a new perimeter method to the Shape interface to calculate
   the perime‐ ter of a shape, and implementing the method for
   Circle and Rectangle:
*/

package main

import "fmt"

func main() {
	c := &circle{radius: 10}
	r := &rectangle{length: 10, breadth: 20}
	fmt.Println("Perimeter of  circle =", getPerimeter(c))
	fmt.Println("Perimeter of rectangle =", getPerimeter(r))
}

type shapes interface {
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c *circle) perimeter() float64 {
	return 2 * (3.14) * c.radius
}

type rectangle struct {
	length, breadth float64
}

func (r *rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func getPerimeter(s shapes) float64 {
	return s.perimeter()
}
