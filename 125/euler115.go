// Counting block combinations II
//
// Completed on Fri, 3 Jun 2022, 16:27

package main

import "fmt"

const base = 50
const limit = 1000000

var combinations map[int]int

func solve(n int) (res int) {
	var right, left int

	r, ok := combinations[n]
	if ok {
		return r
	}

	res = 1
	for right = n - 1; right > (base - 2); right-- {
		for left = 0; right > (base + left - 2); left++ {
			if n-right > (base + 1) {
				res += solve(n - right - 2)
			} else {
				res++
			}
		}
	}

	combinations[n] = res

	return
}

func main() {
	combinations = make(map[int]int)

	var length = base
	for {
		if solve(length) > limit {
			fmt.Println(length)
			break
		}

		length++
	}
}
