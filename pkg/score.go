package yahtzee

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
	one   uint8
	two   uint8
	three uint8
	four  uint8
	five  uint8
	six   uint8
	bonus bool
	total uint8
	mask  [6]bool
}

type LowerSection struct {
	tKind     uint8
	fKind     uint8
	fHouse    bool
	sStraight bool
	lStraight bool
	yahtzee   bool
	chance    uint8
	yBonus    uint16
	total     uint16
	mask      [7]bool
}
