// Totient permutation
//
// Completed on Tue, 17 Jan 2017, 01:30

package main

import (
	"fmt"
	"github.com/morontt/projecteuler/goeuler"
	"math"
	"runtime"
	"sort"
	"strings"
)

var (
	primes   []uint64
	maxprime uint64
	ncpu     uint64
)

const nmax uint64 = 10000000

type IterationResult struct {
	Min  float64
	Find uint64
}

func totient(x uint64) uint64 {
	var (
		limit, res uint64
		divs       map[uint64]uint
	)

	limit = 1 + uint64(math.Sqrt(float64(x)))
	divs = make(map[uint64]uint)
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
				limit = 1 + uint64(math.Sqrt(float64(x)))
				flag = true
				break
			}
		}
		if !flag {
			divs[x] += 1
		}
	}

	res = 1
	for idx, z := range divs {
		res *= uint64(math.Pow(float64(idx), float64(z-1))) * (idx - 1)
	}

	return res
}

func isPermutation(x, y uint64) bool {
	sx := strings.Split(fmt.Sprintf("%d", x), "")
	sy := strings.Split(fmt.Sprintf("%d", y), "")

	res := false
	if len(sx) == len(sy) {
		sort.Sort(sort.StringSlice(sx))
		sort.Sort(sort.StringSlice(sy))
		res = strings.Join(sx, "") == strings.Join(sy, "")
	}

	return res
}

func number_cycle(start uint64, results chan<- IterationResult) {
	var (
		min  float64
		find uint64
	)

	min = 2.0
	for i := start; i < nmax; i += ncpu {
		j := totient(i)
		if isPermutation(i, j) {
			m := float64(i) / float64(j)
			if m < min {
				min = m
				find = i
			}
		}
	}

	results <- IterationResult{min, find}
}

func main() {
	var (
		min  float64
		find uint64
		res  IterationResult
	)

	ncpu = uint64(runtime.NumCPU())
	maxprime = 1 + uint64(math.Sqrt(float64(nmax)))
	primes = append(primes, 2, 3)

	j := uint64(0)
	for i := uint64(0); j < maxprime; i++ {
		j = 6*i + 1
		if goeuler.IsPrime(j) {
			primes = append(primes, j)
		}

		j = j + 4
		if goeuler.IsPrime(j) {
			primes = append(primes, j)
		}
	}

	results := make(chan IterationResult)

	for i := uint64(0); i < ncpu; i++ {
		go number_cycle(i+2, results)
	}

	min = 2.0
	for k := uint64(0); k < ncpu; k++ {
		res = <-results
		if res.Min < min {
			min = res.Min
			find = res.Find
		}
	}

	fmt.Println(find)
}
