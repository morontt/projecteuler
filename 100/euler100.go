// Arranged Probability
//
// Completed on Sat, 9 Mar 2024, 20:42

package main

import (
	"fmt"
	"math/big"
)

func main() {
	i := big.NewInt(1000_000_000_000)
	e := big.NewInt(1)

	for {
		if checkDiscs(i) {
			break
		}
		i.Add(i, e)
	}
}

func checkDiscs(m *big.Int) (res bool) {
	var lb, rb, s, s0, s01, s1, s11 *big.Int

	one := big.NewInt(1)
	two := big.NewInt(2)

	lb = big.NewInt(1)
	rb = new(big.Int).Set(m)
	s0 = new(big.Int).Set(m)
	s01 = new(big.Int).Set(m)
	s01.Sub(s01, one)
	s0.Mul(s0, s01)

	s = new(big.Int)
	s1 = new(big.Int)
	s11 = new(big.Int)

	lbt := new(big.Int)

	for {
		lbt.Set(lb)
		lbt.Add(lbt, one)
		if lbt.Cmp(rb) == 0 || lb.Cmp(rb) == 0 {
			break
		}

		s.Set(lb)
		s.Add(s, rb)
		s.Div(s, two)
		s1.Set(s)
		s11.Set(s)
		s11.Sub(s11, one)
		s1.Mul(s1, s11)
		s1.Mul(s1, two)

		cmp := s0.Cmp(s1)
		if cmp == 0 {
			fmt.Printf("> %d / %d\n", s, m)
			res = true

			break
		}

		if cmp < 0 {
			rb.Set(s)
		} else {
			lb.Set(s)
		}
	}

	return
}
