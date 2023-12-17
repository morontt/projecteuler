// Prime Power Triples
//
// Completed on Sun, 17 Dec 2023, 21:38

package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"

	"github.com/morontt/projecteuler/goeuler"
)

const nmax = 50000000

type trCache struct {
	mu   sync.RWMutex
	data map[int]bool
}

var (
	primes       []int
	fourth       map[int]bool
	triplesCache trCache
	ncpu         int
)

func main() {
	var j, k, res int

	maxprime := int(math.Sqrt(float64(nmax)))
	capacity := int(0.5 * math.Sqrt(float64(nmax)))

	primes = make([]int, 0, capacity)
	primes = append(primes, 2, 3)

	fourth = make(map[int]bool)
	fourth[16] = true
	fourth[81] = true

	triplesCache.data = make(map[int]bool)

	for i := 0; j < maxprime; i++ {
		j = 6*i + 1
		if goeuler.IsPrime(j) {
			if j < maxprime {
				primes = append(primes, j)
				k = j * j
				k *= k
				if k < nmax {
					fourth[k] = true
				}
			}
		}

		j = j + 4
		if goeuler.IsPrime(j) {
			if j < maxprime {
				primes = append(primes, j)
				k = j * j
				k *= k
				if k < nmax {
					fourth[k] = true
				}
			}
		}
	}

	ncpu = runtime.NumCPU()
	results := make(chan int)

	for i := 0; i < ncpu; i++ {
		go checkNumbers(i, results)
	}

	for i := 0; i < ncpu; i++ {
		res += <-results
	}

	fmt.Printf("> %d", res)
}

func checkNumbers(start int, results chan<- int) {
	var (
		p, delta, result int
	)

	for i := start; i <= nmax; i += ncpu {
		for _, p = range primes {
			delta = i - p*p
			if delta < 0 {
				break
			}

			if checkTriples(delta) {
				result += 1
				break
			}
		}
	}

	results <- result
}

func checkTriples(x int) bool {
	triplesCache.mu.RLock()
	res, found := triplesCache.data[x]
	triplesCache.mu.RUnlock()
	if found {
		return res
	}

	for _, p := range primes {
		y := x - p*p*p
		if y < 0 {
			break
		}

		if _, ok := fourth[y]; ok {
			triplesCache.mu.Lock()
			triplesCache.data[x] = true
			triplesCache.mu.Unlock()

			return true
		}
	}

	triplesCache.mu.Lock()
	triplesCache.data[x] = false
	triplesCache.mu.Unlock()

	return false
}
