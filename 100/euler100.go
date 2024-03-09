// Arranged Probability
//
// Completed on

// > 112529341 / 159140520

package main

import (
	"fmt"
)

func main() {
	var i uint64 = 100_000_000

	for {
		if checkDiscs(i) {
			break
		}
		i++
	}
}

func checkDiscs(m uint64) (res bool) {
	var lb, rb, s, s0, s1 uint64

	lb = 1
	rb = m
	s0 = m * (m - 1)
	for {
		if lb == rb || lb+1 == rb {
			break
		}
		s = (lb + rb) >> 1
		s1 = 2 * s * (s - 1)

		if s0 == s1 {
			fmt.Printf("> %d / %d\n", s, m)
			res = true

			break
		}

		if s0 < s1 {
			rb = s
		} else {
			lb = s
		}
	}

	return
}
