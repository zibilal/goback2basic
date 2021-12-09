package presentation

import (
	"fmt"
	"testing"
)

type Lagu struct {
	ID string
	Artist string
}

func (l Lagu) PrintArtist() {
	fmt.Println(l.Artist)
}

func TestLagu(t *testing.T) {
	llagu := new(Lagu)
	llagu.ALagu()
}

func (l *Lagu) ALagu() {
	l.ID = l.ID + "-opened"
	fmt.Println("Lagu: ", l)
	fmt.Println("Artist: ", l.Artist)
}
