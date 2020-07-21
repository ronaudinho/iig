package main

import (
	"github.com/ronaudinho/iig"
)

func main() {
	bus := iig.NewBusquets()
	pir := iig.NewPirlo()
	nas := iig.NewNash()

	// using *iig.Busquets
	play(bus)
	dive(bus)

	// defend() will fail because *iig.Busquets does not implement Defend()
	// to fix this, simply add Defend() method on *iig.Busquets()
	// i can probably use flag to
	defend(bus)

	// ---------

	// using *iig.Nash
	// play(nas) succeeds because *iig.Nash has PlayFootball()
	play(nas)
	pass(nas)
	ismultitalented(nas)

	// ---------

	// comparing *iig.Busquets, *iig.Pirlo, and *iig.Nash
	// different function implementation for Pass()
	// but still works because all implements Footballer interface
	pass(pir)
	pass(bus)
	pass(nas)

}

/*
// want to play busquets
func play(b *iig.Busquets) {
	b.PlayFootball()
}

// want to play pirlo as well, how?
func play(p *iig.Pirlo) {
	p.PlayFootball()
}
*/

// ---------

// simply define Footballer interface here
// anyone that can PlayFootball() can play()
type Footballer interface {
	PlayFootball()
}

func play(f Footballer) {
	f.PlayFootball()
}

// ---------

func defend(d Defender) {
	d.Defend()
}

// ---------
// multitude of other interfaces used as examples

func pass(p Passer) {
	p.Pass()
}

func dive(d Diver) {
	d.Dive()
}

type Basketballer interface {
	PlayBasketball()
}

type Passer interface {
	Pass()
}

type Defender interface {
	Defend()
}

type Diver interface {
	Dive()
}

// ---------

// composition of interface
// must be footballer and basketballer
type MultiTalent interface {
	Footballer
	Basketballer
	Play()
}

// this achieves the same effect as MultiTalent interface above
type MultiTalentByFunc interface {
	PlayFootball()
	PlayBasketball()
	Play()
}

func ismultitalented(m MultiTalent) {
	m.PlayFootball()
	m.PlayBasketball()
	m.Play()
}
