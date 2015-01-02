// Pandigital products
//
// Completed on Fri, 2 Jan 2015, 22:26

package main

import (
    "fmt"
    "strconv"
)

func check_uniq (z uint64, nums *map[byte]bool) bool {
    zs := strconv.FormatUint(z, 10)
    z_length := len(zs)

    used_numbers := *nums

    for i := 0; i < z_length; i++ {
        if !used_numbers[zs[i]] {
            used_numbers[zs[i]] = true
        } else {
            return false
        }
    }

    *nums = used_numbers

    return true
}

func check_pandiag (a, b uint64, length int) (bool, uint64) {
    c := a * b
    cs := strconv.FormatUint(c, 10)
    prod_length := len(cs)

    if prod_length != length {
        return false, c
    }

    used_numbers := map[byte]bool{
        48: true,  //byte for "0"
        49: false, //byte for "1" etc.
        50: false,
        51: false,
        52: false,
        53: false,
        54: false,
        55: false,
        56: false,
        57: false,
    }

    if !check_uniq(a, &used_numbers) {
        return false, c
    }

    if !check_uniq(b, &used_numbers) {
        return false, c
    }

    if !check_uniq(c, &used_numbers) {
        return false, c
    }

    return true, c
}

func main() {
    var (
        i, j, sum uint64
    )

    elements := make(map[uint64]bool)
    sum = 0

    for i = 12; i < 99; i++ {
        for j = 123; j < 988; j++ {
            ok, prod := check_pandiag(i, j, 4)
            if ok {
                _, test := elements[prod]
                if !test {
                    elements[prod] = true
                    sum += prod
                    fmt.Println(prod)
                }
            }
        }
    }

    for i = 2; i < 10; i++ {
        for j = 1234; j < 9877; j++ {
            ok, prod := check_pandiag(i, j, 4)
            if ok {
                _, test := elements[prod]
                if !test {
                    elements[prod] = true
                    sum += prod
                    fmt.Println(prod)
                }
            }
        }
    }

    fmt.Println("----")
    fmt.Println(sum)
}
