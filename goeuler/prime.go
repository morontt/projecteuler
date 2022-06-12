package goeuler

import "math"

func IsPrime(x uint64) (flag bool) {
	var (
		limit uint64
	)

	if x == 1 || x == 4 {
		return false
	}

	if x == 2 || x == 3 || x == 5 {
		return true
	}

	if x%2 == 0 {
		return false
	}

	if x%3 == 0 {
		return false
	}

	if x%5 == 0 {
		return false
	}

	limit = 1 + uint64(math.Sqrt(float64(x)))
	flag = true

	for i := uint64(6); i < limit; i += 6 {
		if x%(i+1) == 0 {
			flag = false
			break
		}

		if x%(i+5) == 0 {
			flag = false
			break
		}
	}

	return
}
