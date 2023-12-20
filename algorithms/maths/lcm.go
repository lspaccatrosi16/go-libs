package maths

func Lcm[T MInt](nums ...T) T {
	if len(nums) < 1 {
		return 0
	}
	r := nums[0]

	for i := 1; i < len(nums); i++ {
		r = ((nums[i] * r) / Hcf(nums[i], r))
	}

	return r
}
