package yahtzee

import (
	"errors"
)

type Card struct {
	Name   string
	Rounds [6]Round
	Total  uint16
}

type Round struct {
	u     *UpperSection
	l     *LowerSection
	total uint16
}

type UpperSection struct {
	scores []uint8
	bonus  bool
	total  uint8
	mask   [6]bool
}

type LowerSection struct {
	tKind     uint8
	fKind     uint8
	fHouse    bool
	sStraight bool
	lStraight bool
	yahtzee   bool
	chance    uint8
	yBonus    uint8
	total     uint16
	mask      [7]bool
}

type Section interface {
	UpdateSetion()
	pickScore(i int) error
	checkBitmask(i int)
}

func (s *UpperSection) Update() {
	// i := 0
	// s.pickScore(i)
}

func (s *LowerSection) Update() {

}

func (s *UpperSection) pickScore(i int, n uint8) {
	if i < 0 || i > 5 {
		panic(errors.New("Index not in range"))
	}

	s.scores[i] = n
	s.mask[i] = true
}

func (s *LowerSection) pickScore(i int, n uint8) error {
	var err error = nil
	switch i {
	case 0:
		s.tKind = n
		break
	case 1:
		s.fKind = n
		break
	case 2:
		s.fHouse = true
		break
	case 3:
		s.sStraight = true
		break
	case 4:
		s.lStraight = true
		break
	case 5:
		s.yahtzee = true
		break
	case 6:
		s.chance = n
		break
	case 7:
		if s.yBonus == 3 {
			return errors.New("Cannot have more than three bonus yahtzees")
		}
		s.yBonus++
	default:
		return errors.New("Index not in range")
	}

	if err == nil {
		s.mask[i] = true
	}
	return err
}

func (s *UpperSection) checkBitmask(i int) bool {
	return s.mask[i]
}

func (s *LowerSection) checkBitmask(i int) bool {
	return s.mask[i]
}
