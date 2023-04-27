// Counting fractions
//
// Completed on Thu, 27 Apr 2023, 16:52

package main

import (
	"fmt"
	"math"
	"runtime"

	"github.com/morontt/projecteuler/goeuler"
)

const nmax = 1000000

var (
	primes   []uint
	maxprime uint
	ncpu     uint
)

func divisors(x uint) map[uint]uint {
	limit := 1 + uint(math.Sqrt(float64(x)))
	divs := make(map[uint]uint)
	flag := true

	for flag {
		flag = false
		for _, p := range primes {
			if p > limit {
				break
			}

			if x%p == 0 {
				divs[p] += 1
				x /= p
				limit = 1 + uint(math.Sqrt(float64(x)))
				flag = true
				break
			}
		}

		if !flag && x != 1 {
			divs[x] += 1
		}
	}

	return divs
}

func fractions(start uint, results chan<- uint) {
	var (
		flag   bool
		result uint
	)

	for i := start; i <= nmax; i += ncpu {
		result += 1
		divs := divisors(i)
		for j := uint(2); j < i; j++ {
			flag = true
			for idx := range divs {
				if j%idx == 0 {
					flag = false
					break
				}
			}

			if flag {
				result += 1
			}
		}
	}

	results <- result
}

func main() {
	var res uint

	ncpu = uint(runtime.NumCPU())
	maxprime = 1 + uint(math.Sqrt(float64(nmax)))
	primes = append(primes, 2, 3)

	j := uint(0)
	for i := uint(0); j < maxprime; i++ {
		j = 6*i + 1
		if goeuler.IsPrime(j) {
			primes = append(primes, j)
		}

		j = j + 4
		if goeuler.IsPrime(j) {
			primes = append(primes, j)
		}
	}

	fmt.Printf("%v\n", primes)

	results := make(chan uint)

	for i := uint(0); i < ncpu; i++ {
		go fractions(i+2, results)
	}

	for k := uint(0); k < ncpu; k++ {
		res += <-results
	}

	fmt.Printf("%v\n", res)
}
