package main

import "fmt"

func rollNorth(arr [][]rune) {

	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j][i] == 'O' {
				pos := j
				for k := j - 1; k >= 0; k-- {
					if arr[k][i] == '.' {
						pos = k
					} else {
						break
					}
				}
				arr[pos][i], arr[j][i] = arr[j][i], arr[pos][i]
			}
		}
	}
}

func rollSouth(arr [][]rune) {

	for i := 0; i < len(arr[0]); i++ {
		for j := len(arr) - 1; j >= 0; j-- {
			if arr[j][i] == 'O' {
				pos := j
				for k := j + 1; k < len(arr); k++ {
					if arr[k][i] == '.' {
						pos = k
					} else {
						break
					}
				}
				arr[pos][i], arr[j][i] = arr[j][i], arr[pos][i]
			}
		}
	}
}

func rollWest(arr [][]rune) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 'O' {
				pos := j
				for k := j - 1; k >= 0; k-- {
					if arr[i][k] == '.' {
						pos = k
					} else {
						break
					}
				}
				arr[i][pos], arr[i][j] = arr[i][j], arr[i][pos]
			}
		}
	}
}

func rollEast(arr [][]rune) {

	for i := 0; i < len(arr); i++ {
		for j := len(arr[0]) - 1; j >= 0; j-- {
			if arr[i][j] == 'O' {
				pos := j
				for k := j + 1; k < len(arr[0]); k++ {
					if arr[i][k] == '.' {
						pos = k
					} else {
						break
					}
				}
				arr[i][pos], arr[i][j] = arr[i][j], arr[i][pos]
			}
		}
	}
}

func solve(arr [][]rune) int {

	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 'O' {
				ans += len(arr) - i
			}
		}
	}
	return ans
}

func convert(arr [][]rune) string {

	ans := ""
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			ans += string(rune(arr[i][j]))
		}
	}
	return ans
}

func cycle(arr [][]rune) {
	rollNorth(arr)
	rollWest(arr)
	rollSouth(arr)
	rollEast(arr)
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

	rollNorth(arr)
	return solve(arr)
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

	mp := make(map[string]int)

	for i := 1; i <= 1000000000; i++ {
		cycle(arr)
		if elem, ok := mp[convert(arr)]; ok {
			rem := (1000000000 - elem) % (i - elem)
			for j := 0; j < rem; j++ {
				cycle(arr)
			}
			return solve(arr)
		}
		mp[convert(arr)] = i

	}
	return 1
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
