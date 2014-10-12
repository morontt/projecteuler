// Coded triangle numbers
//
// Completed on Sun, 12 Oct 2014, 22:04

package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
)

func alpha (str string) uint {
    var (
        res uint = 0
    )

    for _, letter := range str {
        res += uint(letter - 64)
    }

    return res
}

func main() {
    var (
        data   []string
        result uint = 0
        i, j   uint
    )

    file, _ := os.Open("p042_words.txt")

    defer file.Close()

    reader := csv.NewReader(file)
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else {
            data = record
        }
    }

    triangles := make(map[uint]bool)
    //max alpha(...) = 192
    for i = 0; i < 20; i++ {
        j = (i * i + i) / 2
        triangles[j] = true
    }

    for _, val := range data {
        _, ok := triangles[alpha(val)]
        if ok {
            result++
        }
    }

    fmt.Println(result)
}
