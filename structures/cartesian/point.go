package cartesian

import "fmt"

type Coordinate [2]int

func (c Coordinate) Transform(xoffset, yoffset int) Coordinate {
	return Coordinate{c[0] + xoffset, c[1] + yoffset}
}

func (c Coordinate) Resultant(c1 Coordinate) Coordinate {
	return Coordinate{c[0] + c1[0], c[1] + c1[1]}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c[0], c[1])
}
