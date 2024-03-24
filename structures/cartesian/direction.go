package cartesian

import (
	"github.com/lspaccatrosi16/go-libs/structures/enum"
)

type Direction int

const (
	North Direction = iota
	NorthEast
	East
	SouthEast
	South
	SouthWest
	West
	NorthWest
	NoDirection
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case NorthEast:
		return "NorthEast"
	case East:
		return "East"
	case SouthEast:
		return "SouthEast"
	case South:
		return "South"
	case SouthWest:
		return "SouthWest"
	case West:
		return "West"
	case NorthWest:
		return "NorthWest"
	case NoDirection:
		return "Auto"
	default:
		return enum.InvalidString
	}
}

func (s Direction) FromI(i int) enum.Enum {
	return Direction(i)
}

func (s Direction) IsValid() bool {
	return s.String() != enum.InvalidString
}

func (s Direction) Coordinates() Coordinate {
	switch s {
	case North:
		return Coordinate{0, -1}
	case NorthEast:
		return Coordinate{1, -1}
	case East:
		return Coordinate{1, 0}
	case SouthEast:
		return Coordinate{1, 1}
	case South:
		return Coordinate{0, 1}
	case SouthWest:
		return Coordinate{-1, 1}
	case West:
		return Coordinate{-1, 0}
	case NorthWest:
		return Coordinate{-1, -1}
	default:
		return Coordinate{0, 0}
	}
}

func (s Direction) NumberCw(d Direction) int {
	return (int(d) - int(s) + 8) % 8
}

func (s Direction) NumberAcw(d Direction) int {
	return (int(s) - int(d) + 8) % 8
}

func (s Direction) NextCw() Direction {
	if s == NorthWest {
		return North
	} else {
		return Direction(int(s) + 1)
	}
}

func (s Direction) NextAcw() Direction {
	if s == North {
		return NorthWest
	} else {
		return Direction(int(s) - 1)
	}
}

func (s Direction) Opposite() Direction {
	if s == NoDirection {
		return NoDirection
	}
	return Direction((int(s) + 4) % 8)
}

func (s Direction) FromCoordinates(c Coordinate) Direction {
	switch {
	case c[0] > 0 && c[1] == 0:
		return East
	case c[0] < 0 && c[1] == 0:
		return West
	case c[0] == 0 && c[1] > 0:
		return South
	case c[0] == 0 && c[1] < 0:
		return North
	case c[0] > 0 && c[1] > 0:
		return SouthEast
	case c[0] > 0 && c[1] < 0:
		return NorthEast
	case c[0] < 0 && c[1] > 0:
		return SouthWest
	case c[0] < 0 && c[1] < 0:
		return NorthWest
	default:
		return NoDirection
	}
}

func (s Direction) FromXy(x, y int) Direction {
	return s.FromCoordinates(Coordinate{x, y})
}

func CardinalPositions() []Direction {
	return []Direction{North, East, South, West, NorthEast, SouthEast, SouthWest, NorthWest}
}
