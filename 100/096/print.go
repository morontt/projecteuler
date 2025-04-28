package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printField(f field) {
	var (
		// https://symbl.cc/en/unicode-table/#box-drawing
		tt = []string{
			string(rune(9484)), // 0. down-right
			string(rune(9472)), // 1. horizontal
			string(rune(9488)), // 2. left-down
			string(rune(9516)), // 3. left-down-right
			string(rune(9500)), // 4. up-right-down
			string(rune(9532)), // 5. cross
			string(rune(9474)), // 6. vertical
			string(rune(9508)), // 7. up-left-down
			string(rune(9492)), // 8. up-right
			string(rune(9524)), // 9. left-up-right
			string(rune(9496)), // 10. left-up
		}
		line string
	)

	line = tt[0] + strings.Repeat(tt[1], 9) + tt[3] + strings.Repeat(tt[1], 9) + tt[3] + strings.Repeat(tt[1], 9) + tt[2]
	fmt.Println(line)

	for j := range 3 {
		for i := range 3 {
			line = tt[6] + f[3*j+i][0].s() + f[3*j+i][1].s() + f[3*j+i][2].s() +
				tt[6] + f[3*j+i][3].s() + f[3*j+i][4].s() + f[3*j+i][5].s() +
				tt[6] + f[3*j+i][6].s() + f[3*j+i][7].s() + f[3*j+i][8].s() + tt[6]

			fmt.Println(line)
		}

		if j != 2 {
			line = tt[4] + strings.Repeat(tt[1], 9) + tt[5] + strings.Repeat(tt[1], 9) + tt[5] + strings.Repeat(tt[1], 9) + tt[7]
			fmt.Println(line)
		}
	}

	line = tt[8] + strings.Repeat(tt[1], 9) + tt[9] + strings.Repeat(tt[1], 9) + tt[9] + strings.Repeat(tt[1], 9) + tt[10]
	fmt.Println(line)
}

func (c *cell) s() string {
	if c.value == 0 {
		return " . "
	}

	return " " + strconv.Itoa(c.value) + " "
}

func (c *cell) sBlock() string {
	var test int = 9

	if c.blocked[test] {
		return " + "
	}

	if c.value == test {
		return " " + strconv.Itoa(test) + " "
	}

	return " . "
}
