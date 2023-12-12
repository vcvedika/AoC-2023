package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func position(a int, b []int) int {

	ans := -1
	for i, elem := range b {
		if elem < a {
			ans = i
		}
	}
	return ans + 1
}

func part1() int {

	var arr [][]rune
	var exprow []int
	var expcol []int
	k := 0

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		expand_row := true
		var a []rune
		for _, elem := range line {
			a = append(a, rune(elem))
			if elem == '#' {
				expand_row = false
			}
		}
		arr = append(arr, a)
		if expand_row {
			exprow = append(exprow, k)
		}
		k++
	}

	row := len(arr)
	col := len(arr[0])

	for i := 0; i < col; i++ {
		expand_col := true
		for j := 0; j < row; j++ {
			if arr[j][i] == '#' {
				expand_col = false
			}
		}
		if expand_col {
			expcol = append(expcol, i)
		}
	}

	ans := 0.0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if arr[i][j] == '#' {
				for k := 0; k < row; k++ {
					for l := 0; l < col; l++ {
						if arr[k][l] == '#' {
							temp_row := position(k, exprow)
							temp_col := position(l, expcol)
							tempr_init := position(i, exprow)
							tempc_init := position(j, expcol)

							ans += math.Abs(float64(temp_row+k-i-tempr_init)) + math.Abs(float64(temp_col+l-j-tempc_init))
						}
					}
				}
			}
		}
	}

	return int(ans) / 2
}

func part2() int {

	var arr [][]rune
	var exprow []int
	var expcol []int
	k := 0

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		expand_row := true
		var a []rune
		for _, elem := range line {
			a = append(a, rune(elem))
			if elem == '#' {
				expand_row = false
			}
		}
		arr = append(arr, a)
		if expand_row {
			exprow = append(exprow, k)
		}
		k++
	}

	row := len(arr)
	col := len(arr[0])

	for i := 0; i < col; i++ {
		expand_col := true
		for j := 0; j < row; j++ {
			if arr[j][i] == '#' {
				expand_col = false
			}
		}
		if expand_col {
			expcol = append(expcol, i)
		}
	}

	ans := 0.0
	val := 1000000 - 1
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if arr[i][j] == '#' {
				for k := 0; k < row; k++ {
					for l := 0; l < col; l++ {
						if arr[k][l] == '#' {
							temp_row := position(k, exprow) * val
							temp_col := position(l, expcol) * val
							tempr_init := position(i, exprow) * val
							tempc_init := position(j, expcol) * val

							ans += math.Abs(float64(temp_row+k-i-tempr_init)) + math.Abs(float64(temp_col+l-j-tempc_init))
						}
					}
				}
			}
		}
	}

	return int(ans) / 2
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
