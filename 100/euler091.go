// Right Triangles with Integer Coordinates
//
// Completed on Wed, 27 Dec 2023, 19:34

package main

import (
	"fmt"
)

const (
	n = 51
	l = n * n
)

func main() {
	var (
		i, xa, ya       int
		j, xb, yb       int
		res, s1, s2, s3 int
	)

	for i = 1; i < l-1; i++ {
		xa = i % n
		ya = i / n
		for j = i + 1; j < l; j++ {
			xb = j % n
			yb = j / n

			s1 = xa*xb + ya*yb
			if s1 == 0 {
				res += 1

				continue
			}

			s2 = xa*(xb-xa) + ya*(yb-ya)
			if s2 == 0 {
				res += 1

				continue
			}

			s3 = xb*(xb-xa) + yb*(yb-ya)
			if s3 == 0 {
				res += 1
			}
		}
	}

	fmt.Printf("> %d", res)
}
