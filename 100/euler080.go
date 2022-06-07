// Square root digital expansion
//
// Completed on Tue, 7 Jun 2022, 23:13

package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const (
	prec    = 500
	epsilon = 200
)

var iterations int

func bigSqrt(y int64) big.Float {
	square := new(big.Float).SetPrec(prec).SetInt64(y)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	a := new(big.Float).SetPrec(prec).SetInt64(1)
	b := new(big.Float).SetPrec(prec).SetInt64(y)

	t := new(big.Float)
	ft := new(big.Float)

	for i := 0; i < iterations; i++ {
		t.Add(a, b)
		t.Mul(t, half)
		ft.Mul(t, t)

		if ft.Cmp(square) == 1 {
			b.Copy(t)
		} else {
			a.Copy(t)
		}
	}

	fmt.Printf("sqrt(%d) = %.40f\n", y, t)

	return *t
}

func sumDigits(z int64) int {
	var sum int

	f := bigSqrt(z)
	str := strings.Split(f.Text('g', 101), "")
	for i := 0; i <= 100; i++ {
		if i == 1 {
			continue
		}

		z, _ := strconv.Atoi(str[i])
		sum += z
	}

	return sum
}

func main() {
	var (
		squares map[int64]bool
		i       int64
		sum     int
	)

	iterations = int((epsilon + 2) * math.Log2(10.0))

	squares = make(map[int64]bool)
	for i = 1; i <= 11; i++ {
		squares[i*i] = true
	}

	for i = 1; i < 100; i++ {
		_, ok := squares[i]
		if !ok {
			sum += sumDigits(i)
		}
	}

	fmt.Println(sum)
}
