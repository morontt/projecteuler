// Largest prime factor
//
// Completed on Sun, 7 Sep 2014, 21:57

package main

import (
    "fmt"
    "math"
)

func div(x uint64) (uint64, bool) {
    var (
        res, limit uint64
    )

    limit = 1 + uint64(math.Sqrt(float64(x)))
    flag := false
    res = x

    for i := uint64(2); i < limit; i++ {
        r := x % i
        if r == 0 {
            res = i
            flag = true
            break
        }
    }

    return res, flag
}

func main() {
    var (
        m, z uint64
        ok   bool = true
    )

    m = 600851475143

    for ok {
        z, ok = div(m)
        m = m / z
    }

    fmt.Println(z)
}
