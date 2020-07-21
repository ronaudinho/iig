package iig

import (
	"fmt"
)

// sample of interface pollution
// creating interface for the sake of it
// instead create interface only when you need it
type Footballer interface {
	PlayFootball()
	TakePenalty()
	TakeFreeKick()
	TakeCornerKick()
}

type FootballerS struct {
	Name string
}

func NewFootballer(nam string) *FootballerS {
	return &FootballerS{Name: nam}
}

func (f *FootballerS) PlayFootball() {
	fmt.Printf("%s is playing football\n", f.Name)
}

func (f *FootballerS) TakePenalty() {
	fmt.Printf("%s takes penalty\n", f.Name)
}

func (f *FootballerS) TakeFreeKick() {
	fmt.Printf("%s takes free kick\n", f.Name)
}

func (f *FootballerS) TakeCornerKick() {
	fmt.Printf("%s takes corner kick\n", f.Name)
}
