package cartesian

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case South:
		return "South"
	case East:
		return "East"
	case West:
		return "West"
	default:
		return "Invalid"
	}
}

func (d Direction) RotateCW() Direction {
	vInt := int(d)
	var newV int

	if vInt == 3 {
		newV = 0
	} else {
		newV = vInt + 1
	}

	return Direction(newV)
}

func (d Direction) RotateACW() Direction {
	vInt := int(d)
	var newV int

	if vInt == 0 {
		newV = 3
	} else {
		newV = vInt - 1
	}

	return Direction(newV)
}
