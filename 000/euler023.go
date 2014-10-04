// Non-abundant sums
//
// Completed on Sat, 4 Oct 2014, 16:43

package main

import (
    "fmt"
    "math"
)

func summDiv(x uint) (uint) {
    var (
        limit, i, res uint
    )

    limit = uint(math.Sqrt(float64(x)))
    res = 1

    for i = 2; i <= limit; i++ {
        r := x % i
        if r == 0 {
            res += i
            res += x / i
        }
    }

    if limit * limit == x {
        res -= limit
    }

    return res
}

func isAbundant(x uint) (bool) {
    if x < 12 {
        return false
    }

    res := false
    if summDiv(x) > x {
        res = true
    }

    return res
}

func main() {
    var (
        i, j, sum, max, limit uint
    )

    limit = 28124

    cache := make(map[uint]bool)
    arr := []uint{}

    for i = 1; i < limit; i++ {
        tmp := isAbundant(i)
        if tmp {
            arr = append(arr, i)
        }
        cache[i] = tmp
    }

    sum = 0

    for i = 1; i < limit; i++ {
        j = 0;
        max = i / 2
        res := true
        for arr[j] <= max {
            k := i - arr[j]
            if cache[k] {
                res = false
            }
            j++
        }

        if res {
            sum += i
        }
    }

    fmt.Println(sum)
}
