package main

import (
	yahtzee "github.com/justincremer/yahtzee/pkg"
)

func main() {
	d := yahtzee.New()
	for i := 0; i < 10; i++ {
		d.Roll(d.KeepMask)
		d.Display()
	}
}
