package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Pair struct {
	first, second int
}

func part1() int {

	ans := 0

	var data [][]rune
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}

		row := make([]rune, 0)
		for _, elem := range input {
			row = append(row, elem)
		}
		data = append(data, row)
	}

	row := len(data)
	col := len(data[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			var number string
			var flag bool

			if unicode.IsDigit(rune(data[i][j])) {
				start_j := j

				for j < col && unicode.IsDigit(rune(data[i][j])) {
					number += string(data[i][j])
					if i+1 < row && data[i+1][j] != '.' && !unicode.IsDigit(rune(data[i+1][j])) {
						flag = true
					}
					if i-1 >= 0 && data[i-1][j] != '.' && !unicode.IsDigit(rune(data[i-1][j])) {
						flag = true
					}
					j++
				}

				last_j := j - 1

				if start_j-1 >= 0 && data[i][start_j-1] != '.' && !unicode.IsDigit(rune(data[i][start_j-1])) {
					flag = true
				}
				if start_j-1 >= 0 && i-1 >= 0 && data[i-1][start_j-1] != '.' && !unicode.IsDigit(rune(data[i-1][start_j-1])) {
					flag = true
				}
				if start_j-1 >= 0 && i+1 < row && data[i+1][start_j-1] != '.' && !unicode.IsDigit(rune(data[i+1][start_j-1])) {
					flag = true
				}

				if last_j+1 < col && data[i][last_j+1] != '.' && !unicode.IsDigit(rune(data[i][last_j+1])) {
					flag = true
				}
				if last_j+1 < col && i-1 >= 0 && data[i-1][last_j+1] != '.' && !unicode.IsDigit(rune(data[i-1][last_j+1])) {
					flag = true
				}
				if last_j+1 < col && i+1 < row && data[i+1][last_j+1] != '.' && !unicode.IsDigit(rune(data[i+1][last_j+1])) {
					flag = true
				}

				if flag {
					val, _ := strconv.Atoi(number)
					ans += val
				}
				j--
			}
		}
	}
	return ans
}

func part2() int {

	ans := 0

	var data [][]rune
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}

		row := make([]rune, 0)
		for _, elem := range input {
			row = append(row, elem)
		}
		data = append(data, row)
	}

	row := len(data)
	col := len(data[0])

	mp := make(map[Pair][]int)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {

			var number string

			if unicode.IsDigit(rune(data[i][j])) {
				start_j := j

				for j < col && unicode.IsDigit(rune(data[i][j])) {
					number += string(data[i][j])
					j++
				}

				val, _ := strconv.Atoi(number)

				j = start_j

				for j < col && unicode.IsDigit(rune(data[i][j])) {
					if i+1 < row && data[i+1][j] == '*' {
						mp[Pair{i + 1, j}] = append(mp[Pair{i + 1, j}], val)
					}
					if i-1 >= 0 && data[i-1][j] == '*' {
						mp[Pair{i - 1, j}] = append(mp[Pair{i - 1, j}], val)
					}
					j++
				}

				last_j := j - 1

				if start_j-1 >= 0 && data[i][start_j-1] == '*' {
					mp[Pair{i, start_j - 1}] = append(mp[Pair{i, start_j - 1}], val)
				}
				if start_j-1 >= 0 && i-1 >= 0 && data[i-1][start_j-1] == '*' {
					mp[Pair{i - 1, start_j - 1}] = append(mp[Pair{i - 1, start_j - 1}], val)
				}
				if start_j-1 >= 0 && i+1 < row && data[i+1][start_j-1] == '*' {
					mp[Pair{i + 1, start_j - 1}] = append(mp[Pair{i + 1, start_j - 1}], val)
				}

				if last_j+1 < col && data[i][last_j+1] == '*' {
					mp[Pair{i, last_j + 1}] = append(mp[Pair{i, last_j + 1}], val)
				}
				if last_j+1 < col && i-1 >= 0 && data[i-1][last_j+1] == '*' {
					mp[Pair{i - 1, last_j + 1}] = append(mp[Pair{i - 1, last_j + 1}], val)
				}
				if last_j+1 < col && i+1 < row && data[i+1][last_j+1] == '*' {
					mp[Pair{i + 1, last_j + 1}] = append(mp[Pair{i + 1, last_j + 1}], val)
				}
				j--
			}
		}
	}

	for _, elem := range mp {
		if len(elem) == 2 {
			ans += elem[0] * elem[1]
		}
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
