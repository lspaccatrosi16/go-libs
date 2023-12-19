package maths

func Pfactor(n int) []int {
	factors := []int{}

	for i := 2; i*i <= n; i++ {
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
