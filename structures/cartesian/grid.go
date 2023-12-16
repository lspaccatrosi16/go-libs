package cartesian

import "fmt"

type CoordinateGrid[T any] map[int]map[int]T

// Add a new coordinate to the grid
func (cg *CoordinateGrid[T]) Add(c Coordinate, val T) {
	if cg == nil {
		cg = new(CoordinateGrid[T])
	}

	if v, ok := (*cg)[c[0]]; ok {
		v[c[1]] = val
	} else {
		(*cg)[0] = map[int]T{c[1]: val}
	}
}

// Swap two points round
func (cg *CoordinateGrid[T]) Swap(c1, c2 Coordinate) {
	if cg == nil {
		return
	}
	(*cg)[c1[0]][c1[1]], (*cg)[c2[0]][c2[1]] = (*cg)[c2[0]][c2[1]], (*cg)[c1[0]][c1[1]]
}

// Get the entry at the coordinate c
func (cg *CoordinateGrid[T]) Get(c Coordinate) T {
	if cg == nil {
		return *new(T)
	}

	return (*cg)[c[0]][c[1]]
}

func (cg *CoordinateGrid[T]) rows() [][]T {
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

func (cg *CoordinateGrid[T]) cols() [][]T {
	if cg == nil {
		return nil
	}

	lines := [][]T{}

	for _, l := range *cg {
		thisCol := []T{}
		for _, c := range l {
			thisCol = append(thisCol, c)
		}
		lines = append(lines, thisCol)
	}

	return lines
}

// Pretty format the grid
func (cg *CoordinateGrid[T]) String() string {
	lines := cg.rows()

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

// Produce a hashable representation of the grid
func (cg *CoordinateGrid[T]) Hash() string {
	arrs := cg.rows()

	hashStr := ""

	for _, l := range arrs {
		for _, v := range l {
			hashStr += fmt.Sprint(v)
		}
	}

	return hashStr
}

// Get the grid's columns
func (cg *CoordinateGrid[T]) GetRows() [][]T {
	return cg.rows()
}

// Get the grid's rows
func (cg *CoordinateGrid[T]) GetCols() [][]T {
	return cg.cols()
}
