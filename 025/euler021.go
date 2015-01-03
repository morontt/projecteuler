// Amicable numbers
//
// Completed on Thu, 18 Sep 2014, 22:43

package main

import (
    "fmt"
    "math"
)

func div(x uint64) (uint64) {
    var (
        limit, i, res uint64
    )

    limit = uint64(math.Sqrt(float64(x)))
    res = 1

    if limit * limit == x {
        res += limit
    }

    for i = 2; i < limit; i++ {
        r := x % i
        if r == 0 {
            res += i
            res += x / i
        }
    }

    return res
}

func main() {
    var (
        n       uint64 = 10000
        s       uint64 = 0
        i, j, k uint64
    )

    for i = 3; i < n; i++ {
        j = div(i)
        if j < i {
            k = div(j)
            if k == i {
                s += i
                s += j
            }
        }
    }

    fmt.Println(s)
}
