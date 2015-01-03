// Pandigital multiples
//
// Completed on Sat, 3 Jan 2015, 11:37

package main

import (
    "fmt"
    "sort"
    "strconv"
)

func is_pandiag(p string) bool {
    p_length := len(p)

    if (p_length != 9) {
        return false
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

    for i := 0; i < p_length; i++ {
        if !used_numbers[p[i]] {
            used_numbers[p[i]] = true
        } else {
            return false
        }
    }

    return true
}

func concatenated_product(x, k int) (string, bool) {
    res := ""
    for i := 1; i <= k; i++ {
        res += strconv.Itoa(x * i)
    }

    return res, (len(res) < 10)
}

func main() {
    var (
        i, k int
        s    string
    )

    pandigits := []string{}

    for i = 2; i < 10; i++ {
        k = 1
        ok := true
        for ok {
            s, ok = concatenated_product(k, i)
            if (is_pandiag(s)) {
                pandigits = append(pandigits, s)
            }
            k += 1
        }
    }

    sort.Strings(pandigits)

    fmt.Println(pandigits[len(pandigits) - 1])
}
