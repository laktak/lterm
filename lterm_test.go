package lterm

import (
	"testing"
)

func Test1(t *testing.T) {

	bg := Bg8(240)
	fg1 := Fg8(235)
	fg2 := Fg8(228)

	Printline(bg, " test", fg1, " foo ", fg2, "bar ", Reset)

}
