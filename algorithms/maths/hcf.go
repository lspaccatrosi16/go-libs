package maths

import (
	"math"
)

func Hcf[T MInt](nums ...T) T {
	var r T
	var highest T = math.MaxInt32

	for i := 0; (i + 1) < len(nums); i++ {
		a := nums[0]
		b := nums[i+1]
		for {
			if a%b == 0 {
				break
			}
			r = a % b
			a = b
			b = r
		}

		if b < highest {
			highest = b
		}
	}

	return highest

}
