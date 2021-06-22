package yahtzee

import (
	"fmt"
	"math/rand"
	"time"
)

type Dice struct {
	Values   [5]uint8
	KeepMask [5]bool
}

func New() *Dice {
	return &Dice{}
}

func (d *Dice) Update(km [5]bool) {
	d.updateMask(km)
	d.roll()
}

func (d *Dice) Display() {
	fmt.Println(d.string())
}

func (d *Dice) string() string {
	v := d.Values
	return fmt.Sprintf("%d %d %d %d %d", v[0], v[1], v[2], v[3], v[4])
}

func (d *Dice) updateMask(km [5]bool) {
	d.KeepMask = km
}

func (d *Dice) roll() {
	rand.Seed(time.Now().Local().UnixNano())
	for i, n := range d.KeepMask {
		if !n {
			time.Sleep(time.Nanosecond)
			d.Values[i] = uint8(rand.Intn(6) + 1)
		}
	}
}
