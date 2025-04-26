// Su Doku
//
// Completed on

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	value   int
	current int
	blocked [10]bool
}

type field [9][9]cell

func main() {
	var ff field

	load(&ff)
	print(ff)
}

func load(f *field) {
	var i, j, v int

	reader := bufio.NewReader(os.Stdin)
	for {
		if j == 9 {
			fmt.Println("done:")
			break
		}

		lineIn, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err: " + err.Error())
			break
		}

		_, err = strconv.Atoi(string(lineIn[0]))
		if err != nil {
			fmt.Println("skip:")
			continue
		} else if len(lineIn) >= 10 {
			for i = 0; i < 10; i++ {
				v, err = strconv.Atoi(string(lineIn[i]))
				if err == nil && v > 0 {
					f.set(j, i, v)
				}
			}
			j++
		}

		fmt.Print("in: " + lineIn)
	}
}

func print(f field) {
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

	for j := 0; j < 3; j++ {
		for i := 0; i < 3; i++ {
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

func (c *cell) block(v int) {
	if !c.blocked[v] {
		c.blocked[v] = true
		c.current += 1
	}
}

func (f *field) set(i, j, v int) {
	if f[i][j].blocked[v] {
		panic(fmt.Sprintf("Invalid cell[%d][%d] value: %d", i, i, v))
	}

	f[i][j].value = v
	for k := 0; k < 9; k++ {
		if k != i {
			f[k][j].block(v)
		}
		if k != j {
			f[i][k].block(v)
		}
	}

	si := i / 3
	sj := j / 3
	for k := 0; k < 3; k++ {
		for m := 0; m < 3; m++ {
			it := 3*si + k
			jt := 3*sj + m
			if it != i && jt != j {
				f[it][jt].block(v)
			}
		}
	}
}
