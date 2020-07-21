package main

import (
	"fmt"

	"github.com/ronaudinho/iig"
)

func main() {
	foo := iig.NewFootballer("Footballer")
	shootpenalty(foo)

	pen := newPenaldo()
	shootpenalty(pen)
}

type PKTaker interface {
	TakePenalty()
}

func shootpenalty(p PKTaker) {
	p.TakePenalty()
}

// surpise penaldo
// penaldo suddenly implements TakePenalty()
// hence penaldo can shootpenalty()
type Penaldo struct {
	Name string
}

func newPenaldo() *Penaldo {
	return &Penaldo{Name: "Penaldo"}
}

func (p *Penaldo) TakePenalty() {
	fmt.Printf("%s misses penalty\n", p.Name)
}
