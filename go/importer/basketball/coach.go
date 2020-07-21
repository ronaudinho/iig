package main

import (
	"github.com/ronaudinho/iig"
)

func main() {
	nas := iig.NewNash()
	play(nas)
	pass(nas)

	bus := iig.NewBusquets()
	play(bus)
}

func play(b Basketballer) {
	b.Play()
}

func pass(p Passer) {
	p.Pass()
}

type Basketballer interface {
	PlayBasketball()
}

type Passer interface {
	Pass()
}
