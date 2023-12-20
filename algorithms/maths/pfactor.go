package maths

func Pfactor[T MInt](n T) []T {
	factors := []T{}

	for i := T(2); i*i <= n; i++ {
		for n%i == 0 {
			n = n / i
			factors = append(factors, i)
		}
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}
