package cartesian

import "fmt"

type CoordinateGrid[T any] map[int]map[int]T

func (cg *CoordinateGrid[T]) Add(c Coordinate, val T) {
	if cg == nil {
		cg = new(CoordinateGrid[T])
	}

	if v, ok := (*cg)[c[0]]; ok {
		v[c[1]] = val
	}

	(*cg)[0] = map[int]T{c[1]: val}
}

func (cg *CoordinateGrid[T]) Swap(c1, c2 Coordinate) {
	if cg == nil {
		return
	}
	(*cg)[c1[0]][c1[1]], (*cg)[c2[0]][c2[1]] = (*cg)[c2[0]][c2[1]], (*cg)[c1[0]][c1[1]]
}

func (cg *CoordinateGrid[T]) Get(c Coordinate) T {
	if cg == nil {
		return *new(T)
	}

	return (*cg)[c[0]][c[1]]
}

func (cg *CoordinateGrid[T]) arrs() [][]T {
	if cg == nil {
		return nil
	}

	maxX := 0
	maxY := 0

	for x, l := range *cg {
		if x > maxX {
			maxX = x
		}
		for y := range l {
			if y > maxY {
				maxY = y
			}
		}
	}

	lines := make([][]T, maxY+1)

	for i := 0; i < len(lines); i++ {
		lines[i] = make([]T, maxX+1)
	}

	for x, l := range *cg {
		for y, v := range l {
			lines[y][x] = v
		}
	}

	return lines
}

func (cg *CoordinateGrid[T]) String() string {
	lines := cg.arrs()

	outStr := ""
	for _, l := range lines {
		s := ""
		for _, v := range l {
			s += fmt.Sprint(v)
		}
		outStr += s + "\n"
	}
	return outStr
}

func (cg *CoordinateGrid[T]) Hash() string {
	arrs := cg.arrs()

	hashStr := ""

	for _, l := range arrs {
		for _, v := range l {
			hashStr += fmt.Sprint(v)
		}
	}

	return hashStr
}
