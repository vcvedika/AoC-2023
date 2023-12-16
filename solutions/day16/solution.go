package main

import "fmt"

type Pair struct {
	i int
	j int
	s string
}

func move(arr [][]rune, temp [][]rune, seen map[Pair]int, i int, j int, direction string) {

	if i < 0 || j < 0 || i == len(arr) || j == len(arr[0]) || seen[Pair{i, j, direction}] == 1 {
		return
	}
	seen[Pair{i, j, direction}] = 1
	temp[i][j] = '#'

	if direction == "up" {
		if arr[i][j] == '.' {
			move(arr, temp, seen, i-1, j, "up")
		} else if arr[i][j] == '-' {
			move(arr, temp, seen, i, j+1, "right")
			move(arr, temp, seen, i, j-1, "left")
		} else if arr[i][j] == '|' {
			move(arr, temp, seen, i-1, j, "up")
		} else if arr[i][j] == '\\' {
			move(arr, temp, seen, i, j-1, "left")
		} else if arr[i][j] == '/' {
			move(arr, temp, seen, i, j+1, "right")
		}
	} else if direction == "down" {
		if arr[i][j] == '.' {
			move(arr, temp, seen, i+1, j, "down")
		} else if arr[i][j] == '-' {
			move(arr, temp, seen, i, j+1, "right")
			move(arr, temp, seen, i, j-1, "left")
		} else if arr[i][j] == '|' {
			move(arr, temp, seen, i+1, j, "down")
		} else if arr[i][j] == '\\' {
			move(arr, temp, seen, i, j+1, "right")
		} else if arr[i][j] == '/' {
			move(arr, temp, seen, i, j-1, "left")
		}
	} else if direction == "left" {
		if arr[i][j] == '.' {
			move(arr, temp, seen, i, j-1, "left")
		} else if arr[i][j] == '-' {
			move(arr, temp, seen, i, j-1, "left")
		} else if arr[i][j] == '|' {
			move(arr, temp, seen, i-1, j, "up")
			move(arr, temp, seen, i+1, j, "down")
		} else if arr[i][j] == '\\' {
			move(arr, temp, seen, i-1, j, "up")
		} else if arr[i][j] == '/' {
			move(arr, temp, seen, i+1, j, "down")
		}
	} else if direction == "right" {
		if arr[i][j] == '.' {
			move(arr, temp, seen, i, j+1, "right")
		} else if arr[i][j] == '-' {
			move(arr, temp, seen, i, j+1, "right")
		} else if arr[i][j] == '|' {
			move(arr, temp, seen, i-1, j, "up")
			move(arr, temp, seen, i+1, j, "down")
		} else if arr[i][j] == '\\' {
			move(arr, temp, seen, i+1, j, "down")
		} else if arr[i][j] == '/' {
			move(arr, temp, seen, i-1, j, "up")
		}
	}
}

func createMatrix(rows, cols int) [][]rune {
	matrix := make([][]rune, rows)
	for i := range matrix {
		matrix[i] = make([]rune, cols)
		for j := range matrix[i] {
			matrix[i][j] = '.'
		}
	}
	return matrix
}

func count(arr [][]rune) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == '#' {
				ans++
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1() int {

	var arr [][]rune
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		arr = append(arr, []rune(input))
	}
	seen := make(map[Pair]int)
	temp := createMatrix(len(arr), len(arr[0]))
	move(arr, temp, seen, 0, 0, "right")

	return count(temp)
}

func part2() int {

	var arr [][]rune
	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}
		arr = append(arr, []rune(input))
	}

	ans := 0
	for i := 0; i < len(arr[0]); i++ {
		seen := make(map[Pair]int)
		temp := createMatrix(len(arr), len(arr[0]))
		move(arr, temp, seen, 0, i, "down")
		ans = max(ans, count(temp))
	}
	for i := 0; i < len(arr[0]); i++ {
		seen := make(map[Pair]int)
		temp := createMatrix(len(arr), len(arr[0]))
		move(arr, temp, seen, len(arr)-1, i, "up")
		ans = max(ans, count(temp))
	}
	for i := 0; i < len(arr); i++ {
		seen := make(map[Pair]int)
		temp := createMatrix(len(arr), len(arr[0]))
		move(arr, temp, seen, i, 0, "right")
		ans = max(ans, count(temp))
	}
	for i := 0; i < len(arr); i++ {
		seen := make(map[Pair]int)
		temp := createMatrix(len(arr), len(arr[0]))
		move(arr, temp, seen, i, len(arr[0])-1, "left")
		ans = max(ans, count(temp))
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
