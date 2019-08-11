// Diophantine equation
//
// Completed on

package main

import (
	"fmt"
	"os"
)

var (
	squires   map[uint64]bool
	maxidx    uint64
	maxsquire uint64
)

func minimalSolution(d uint64) {
	var (
		i, t uint64
	)

	if d == 61 || d == 109 || d == 149 {
		fmt.Printf("%d: skip\n", d)
		return
	}

	i = 1
	for {
		t = d*i*i + 1

		if t > maxsquire {
			fmt.Println("overflow")
			os.Exit(0)
		}

		if _, ok := squires[t]; ok {
			break
		}
		i++
	}

	fmt.Printf("%d: %d - D * %d^2\n", d, t, i)
}

func main() {
	var (
		i uint64
	)

	const dmax = 1000

	maxidx = 200000000
	maxsquire = maxidx * maxidx
	squires = make(map[uint64]bool)
	for i = 1; i <= maxidx; i++ {
		squires[i*i] = true
	}

	for i = 1; i <= dmax; i++ {
		if _, ok := squires[i]; !ok {
			minimalSolution(i)
		}
	}
}
