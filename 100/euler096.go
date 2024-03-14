// Su Doku
//
// Completed on

package main

import (
	"fmt"
	"strings"
)

type cell struct {
	value   int
	current int
	block   [9]bool
}

func main() {
	var field [9][9]cell

	print(field)
}

func print(_ [9][9]cell) {
	// https://symbl.cc/en/unicode-table/#box-drawing
	line0 := string(rune(9484)) + strings.Repeat(string(rune(9472)), 10) + string(rune(9488))

	fmt.Println(line0)
}
