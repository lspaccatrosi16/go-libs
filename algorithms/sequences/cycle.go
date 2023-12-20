package sequences

import (
	"bytes"
	"fmt"
)

type CycleInfo[T comparable] struct {
	Cycle        []T
	RepeatLength int
	StartIdx     int
}

func (c *CycleInfo[T]) String() string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "Repeat Length: %d, Cycle Start: %d\n", c.RepeatLength, c.StartIdx)
	fmt.Fprintln(buf, c.Cycle)
	return buf.String()
}

func (c *CycleInfo[T]) PointAt(n int) int {
	return (n - c.StartIdx + 1) % c.RepeatLength
}

func FindCycle[T comparable](sl []T) CycleInfo[T] {
	info := CycleInfo[T]{}

search:
	for ws := len(sl) / 2; ws > 0; ws-- {
		for p := 0; p <= len(sl)-(2*ws); p++ {
			repSec := sl[p : p+ws]
			before := sl[:p]
			after := sl[p+ws:]

			if existsInSl(before, repSec) && checkCompliance(p, repSec, sl) {
				info.Cycle = repSec
				info.StartIdx = p
				info.RepeatLength = len(repSec)
				break search
			} else if existsInSl(after, repSec) && checkCompliance(p, repSec, sl) {
				info.Cycle = repSec
				info.StartIdx = p
				info.RepeatLength = len(repSec)
				break search
			}
		}
	}

	return info
}

func checkCompliance[T comparable](sp int, sec []T, sl []T) bool {
	endPart := sl[sp:]
	for i := 0; i < len(endPart); i++ {
		if endPart[i] != sec[i%len(sec)] {
			return false
		}
	}

	return true
}

func existsInSl[T comparable](sl []T, sec []T) bool {
	if len(sec) == 0 {
		return false
	}

	if len(sec) > len(sl) {
		return false
	}

	for i := 0; i < len(sl); i++ {
		found := true
		for j := 0; j < len(sec); j++ {
			if i+j >= len(sec) {
				found = false
				break
			}
			if sl[i+j] != sec[j] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}
