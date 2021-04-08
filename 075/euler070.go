// Totient permutation
//
// Completed on Tue, 17 Jan 2017, 01:30

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

var (
	primes   []uint64
	maxprime uint64
)

func isPrime(x uint64) bool {
	var (
		limit uint64
	)

	if x == 1 {
		return false
	}

	if x == 2 || x == 3 {
		return true
	}

	limit = 1 + uint64(math.Sqrt(float64(x)))
	flag := true

	for i := uint64(2); i < limit; i++ {
		if x%i == 0 {
			flag = false
			break
		}
	}

	return flag
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

func main() {
	const nmax = 10000000
	var (
		min  float64
		find uint64
	)

	maxprime = 1 + uint64(math.Sqrt(float64(nmax)))
	primes = append(primes, 2, 3)

	j := uint64(0)
	for i := uint64(0); j < maxprime; i++ {
		j = 6*i + 1
		if isPrime(j) {
			primes = append(primes, j)
		}

		j = j + 4
		if isPrime(j) {
			primes = append(primes, j)
		}
	}

	min = 2.0
	for i := uint64(2); i < nmax; i++ {
		j := totient(i)
		if isPermutation(i, j) {
			m := float64(i) / float64(j)
			if m < min {
				min = m
				find = i
			}
		}
	}

	fmt.Println(find)
}