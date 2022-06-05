// Singular integer right triangles
//
// Completed on Sun, 5 Jun 2022, 12:03

package main

import (
	"fmt"
	"math"
)

const limit = 1500000

var squares map[int]int

func main() {
	var t Triangle
	var triangles map[int]int

	maxSide := limit / 2

	squares = make(map[int]int)
	for i := 1; i <= maxSide; i++ {
		squares[i*i] = i
	}

	triangles = make(map[int]int)
	for i := 2; i < maxSide; i++ {
		jmin := int(float64(i) / math.Sqrt2)
		for j := jmin; j < i; j++ {
			t = Triangle{a: i, b: j}
			if t.valid() {
				cnt, ok := triangles[t.perimetr]
				if ok {
					triangles[t.perimetr] = cnt + 1
				} else {
					triangles[t.perimetr] = 1
				}
			}
		}
	}

	count := 0
	for _, v := range triangles {
		if v == 1 {
			count++
		}
	}

	fmt.Println(count)
}

type Triangle struct {
	a, b, perimetr int
}

func (t *Triangle) valid() bool {
	var c2 int
	if t.a > t.b {
		c2 = (t.a - t.b) * (t.a + t.b)
	} else {
		c2 = (t.b - t.a) * (t.a + t.b)
	}

	c, ok := squares[c2]
	if !ok {
		return false
	}

	if c > t.b {
		return false
	}

	t.perimetr = t.a + t.b + c
	if t.perimetr > limit {
		return false
	}

	return true
}
