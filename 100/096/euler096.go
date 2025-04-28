// Su Doku
//
// Completed on Mon, 28 Apr 2025, 18:22

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	printField(ff)

	move(ff, 0, 0, 0)
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
			for i = range 10 {
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
	for k := range 9 {
		if k != i {
			f[k][j].block(v)
		}
		if k != j {
			f[i][k].block(v)
		}
	}

	si := i / 3
	sj := j / 3
	for k := range 3 {
		for m := range 3 {
			it := 3*si + k
			jt := 3*sj + m
			if it != i && jt != j {
				f[it][jt].block(v)
			}
		}
	}
}

type coord struct {
	x, y int
}

func move(f field, x, y, val int) {
	if val != 0 {
		// fmt.Printf("%d %d: %d\n", x, y, val)
		f.set(x, y, val)
		// printField(f)
	}

	max := 0
	maxCoord := coord{}
	for i := range 9 {
		for j := range 9 {
			if f[i][j].value == 0 && f[i][j].current > max {
				max = f[i][j].current
				maxCoord.x = i
				maxCoord.y = j
			}
		}
	}
	if max == 0 {
		printField(f)
		os.Exit(0)
	}

	if max == 9 {
		return
	}

	var newVal int
	for v := range 9 {
		newVal = v+1
		if busy := f[maxCoord.x][maxCoord.y].blocked[newVal]; !busy {
			move(f, maxCoord.x, maxCoord.y, newVal)
		}
	}
}
