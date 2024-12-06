package main

import "fmt"

/*
Liskov Substitution Principle (LSP)
Objects of a superclass should be replaceable with objects of a subclass without altering the behavior of the program.

In Go:

1) Ensure that any struct implementing an interface can be used interchangeably wherever the interface is used.
2) Stick to the expected behavior defined by the interface.


exmple if there is an interface Bird with method fly if implemented by Peigon it is okay but when implemented by Penguin it breaks the principle
*/

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Println("Area:", s.Area())
}
