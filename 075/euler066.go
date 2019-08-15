// Diophantine equation
//
// Completed on

package main

import (
	"fmt"
	"math"
)

var (
	squires   map[uint64]bool
	maxidx    uint64
	maxsquire uint64
)

type Solve struct {
	d, x, y uint64
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
		if i0 == 4294967295 {
			return Solve{s.d, 0, 4294967295}, true
		}
	}
}

func main() {
	var (
		i        uint64
		s        Solve
		overflow bool
		//arr []Solve
	)

	const dmax = 1000

	maxidx = 50000000
	//maxidx = 1000000
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

	//fmt.Printf("%v\n", arr)
	for _, el := range arr {
		n, _ := otherMinimal(el)
		fmt.Printf("%v\n", n)
	}
}
