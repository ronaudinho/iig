package iig

import (
	"fmt"
)

type Pirlo struct {
	Name    string
	Passing int
}

func NewPirlo() *Pirlo {
	return &Pirlo{Name: "Pirlo", Passing: 90}
}

func (p *Pirlo) Play() {
	fmt.Printf("%s is playing\n", p.Name)
}

func (p *Pirlo) PlayFootball() {
	fmt.Printf("%s is playing football\n", p.Name)
}

func (p *Pirlo) PlayMidfield() {
	fmt.Printf("%s is playing midfield\n", p.Name)
}

func (p *Pirlo) Pass() {
	fmt.Printf("%s passes long\n", p.Name)
}
