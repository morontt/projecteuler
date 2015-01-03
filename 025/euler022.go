// Names scores
//
// Completed on Sat, 20 Sep 2014, 17:16

package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "sort"
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
    )

    file, _ := os.Open("p022_names.txt")

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

    sort.Strings(data)

    for i, val := range data {
        result += alpha(val) * uint(i + 1)
    }

    fmt.Println(result)
}
