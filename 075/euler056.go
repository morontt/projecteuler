// Powerful digit sum
//
// Completed on Sun, 4 Jan 2015, 17:39

package main

import (
    "fmt"
    "math/big"
)

func sumDigit(z *big.Int) int64 {
    var (
        s int64
    )

    s = 0
    str := z.String()

    for i := 0; i < len(str); i++ {
        s += int64(str[i]) - 48
    }

    return s
}

func main() {
    var (
        i, j, s, max int64
    )

    bk := new(big.Int)
    max = 0;

    for i = 2; i < 100; i++ {
        for j = 1; j < 100; j++ {
            bi := big.NewInt(i)
            bj := big.NewInt(j)

            bk.Exp(bi, bj, nil)
            s = sumDigit(bk)

            if max < s {
                max = s
            }

        }
    }

    fmt.Println(max)
}
