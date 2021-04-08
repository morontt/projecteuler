// Diophantine equation
//
// Completed on

package main

import (
	"fmt"
	"math"
	"math/big"
)

var (
	squires   map[uint64]bool
	maxidx    uint64
	maxsquire uint64
)

type Solve struct {
	d, x, y uint64
}

func bigSqrt(y uint64) big.Float {
	const prec = 200
	steps := int(math.Log2(prec))

	square := new(big.Float).SetPrec(prec).SetUint64(y)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	x := new(big.Float).SetPrec(prec).SetInt64(1)
	t := new(big.Float)

	for i := 0; i <= steps; i++ {
		t.Quo(square, x)
		t.Add(x, t)
		x.Mul(half, t)
	}

	fmt.Printf("sqrt(%d) = %.50f\n", y, x)

	return *x
}

func minimalSolution(d uint64) (Solve, bool) {
	var (
		i, t uint64
	)

	i = 1
	for {
		t = d*i*i + 1

		if t > maxsquire {
			fmt.Printf("D: %d overflow\n", d)
			return Solve{d, t, i}, true
		}

		if _, ok := squires[t]; ok {
			break
		}
		i++
	}

	return Solve{d, t, i}, false
}

func otherMinimal(s Solve) (Solve, bool) {
	i0 := s.y
	limit := uint64(4294967295.0 / math.Sqrt(float64(s.d)))

	for {
		t0 := i0 * i0 * s.d
		t1 := t0 + 1

		sqrt := uint64(math.Sqrt(float64(t0)))
		t2 := sqrt * sqrt
		for t2 < t1 {
			sqrt += 1
			t2 = sqrt * sqrt
		}

		if t2 == t1 {
			return Solve{s.d, t2, i0}, false
		}

		i0++
		if i0 == limit {
			fmt.Printf("D: %d overflow (%d)\n", s.d, limit)
			return Solve{s.d, 0, limit}, true
		}
	}
}

func otherBigMinimal(s Solve) {
	var (
		i0, t0, t1, bigOne, bigS big.Int
	)

	i0.SetUint64(s.y)
	bigOne.SetInt64(1)
	bigS.SetUint64(uint64(math.Sqrt(float64(s.d))))

	for {
		t0.Mul(&i0, &i0)
		t0.Add(&t0, &bigOne)

		sqrt := new(big.Int).Set(&i0)
		sqrt.Mul(sqrt, &bigS)
		t1.Mul(sqrt, sqrt)

		for t0.Cmp(&t1) == 1 {
			sqrt.Add(sqrt, &bigOne)
			t1.Mul(sqrt, sqrt)
		}

		if t0.Cmp(&t1) == 0 {
			fmt.Printf("D: %d - %v, %v", s.d, t1, i0)
			return
		}

		i0.Add(&i0, &bigOne)
	}
}

func main() {
	var (
		i        uint64
		s        Solve
		overflow bool
		//arr []Solve
	)

	const dmax = 200

	maxidx = 1000000
	maxsquire = maxidx * maxidx
	squires = make(map[uint64]bool)
	for i = 1; i <= maxidx; i++ {
		squires[i*i] = true
	}

	arr := make([]Solve, 0, 200)

	for i = 1; i <= dmax; i++ {
		if _, ok := squires[i]; !ok {
			s, overflow = minimalSolution(i)
			if overflow {
				//fmt.Printf("%v\n", s)
				arr = append(arr, s)
			}
		}
	}

	arr1 := make([]Solve, 0, 200)

	for _, el := range arr {
		s, overflow = otherMinimal(el)
		if overflow {
			//fmt.Printf("%v\n", s)
			arr1 = append(arr1, s)
		}
	}

	for _, el := range arr1 {
		otherBigMinimal(el)
	}
}
