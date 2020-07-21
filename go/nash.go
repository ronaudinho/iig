package iig

import (
	"fmt"
)

type Nash struct {
	Name    string
	Passing int
}

func NewNash() *Nash {
	return &Nash{Name: "Nash", Passing: 90}
}

func (n *Nash) Play() {
	fmt.Printf("%s is playing\n", n.Name)
}

func (n *Nash) PlayFootball() {
	fmt.Printf("%s is playing football\n", n.Name)
}

func (n *Nash) PlayBasketball() {
	fmt.Printf("%s is playing basketball\n", n.Name)
}

func (n *Nash) PlayPointGuard() {
	fmt.Printf("%s is playing point guard\n", n.Name)
}

func (n *Nash) Pass() {
	fmt.Printf("%s passes alley oop\n", n.Name)
}
