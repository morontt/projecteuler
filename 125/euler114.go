// Counting block combinations I
//
// Completed on Fri, 3 Jun 2022, 11:28

package main

import "fmt"

const length = 50

var combinations map[int]int

func solve(n int) (res int) {
	var right, left int

	r, ok := combinations[n]
	if ok {
		return r
	}

	res = 1
	for right = n - 1; right > 1; right-- {
		for left = 0; right > 1+left; left++ {
			if n-right > 4 {
				res += solve(n - right - 2)
			} else {
				res += 1
			}
		}
	}

	combinations[n] = res

	return
}

func main() {
	combinations = make(map[int]int)

	fmt.Println(solve(length))
}
