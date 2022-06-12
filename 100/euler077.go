// Prime summations
//
// Completed on Sun, 12 Jun 2022, 15:12

package main

import (
	"fmt"
	"github.com/morontt/projecteuler/goeuler"
)

const (
	maxprime = 100
	limit    = 5000
)

var primes []int

func summations(z, idx int) int {
	if z == 0 {
		return 1
	}

	if z == 1 || z < 0 {
		return 0
	}

	var res int
	for i := idx; i >= 0; i-- {
		res += summations(z-primes[i], i)
	}

	return res
}

func main() {
	primes = append(primes, 2, 3)

	j := 0
	for i := 0; j < maxprime; i++ {
		j = 6*i + 1
		if goeuler.IsPrime(uint64(j)) {
			primes = append(primes, j)
		}

		j = j + 4
		if goeuler.IsPrime(uint64(j)) {
			primes = append(primes, j)
		}
	}

	for i := 7; i < maxprime; i++ {
		maxidx := 0
		for key, val := range primes {
			if val <= i {
				maxidx = key
			} else {
				break
			}
		}

		z := summations(i, maxidx)
		if z > limit {
			fmt.Printf("%d	sum: %d\n", i, z)
			break
		}
	}
}
