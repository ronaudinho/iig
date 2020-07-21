package iig

import (
	"fmt"
)

type Busquets struct {
	Name   string
	Diving int
}

func NewBusquets() *Busquets {
	return &Busquets{Name: "Busquets", Diving: 90}
}

func (b *Busquets) Play() {
	fmt.Printf("%s is playing\n", b.Name)
}

func (b *Busquets) PlayFootball() {
	fmt.Printf("%s is playing football\n", b.Name)
}

func (b *Busquets) PlayMidfield() {
	fmt.Printf("%s is playing midfield\n", b.Name)
}

func (b *Busquets) Dive() {
	fmt.Printf("%s dives\n", b.Name)
}

func (b *Busquets) Pass() {
	fmt.Printf("%s passes short\n", b.Name)
}
