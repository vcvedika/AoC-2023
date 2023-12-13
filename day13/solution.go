package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func horizontalMirror(arr [][]rune, extra int) int {

	row := len(arr)
	col := len(arr[0])
	ext := 0

	//horizontal mirror
	for i := 0; i < row-1; i++ {
		//let mirror be present b/w ith and (i+1)th row
		mirror := true
		temp := 0
		ext = 0
		for i+temp+1 <= row-1 && i-temp >= 0 {
			for j := 0; j < col; j++ {
				if rune(arr[i-temp][j]) != rune(arr[i+temp+1][j]) {
					ext++
				}
				if ext > extra {
					mirror = false
					break
				}
			}
			temp++
			if !mirror {
				break
			}
		}
		if mirror && ext == extra {
			return i + 1
		}
	}
	return -1
}

func verticalMirror(arr [][]rune, extra int) int {

	row := len(arr)
	col := len(arr[0])
	ext := 0

	//vertical mirror
	for i := 0; i < col-1; i++ {
		//let mirror be present b/w ith and (i+1)th col
		mirror := true
		temp := 0
		ext = 0
		for i+temp+1 <= col-1 && i-temp >= 0 {
			for j := 0; j < row; j++ {
				if rune(arr[j][i-temp]) != rune(arr[j][i+temp+1]) {
					ext++
				}
				if ext > extra {
					mirror = false
					break
				}
			}
			temp++
			if !mirror {
				break
			}
		}
		if mirror && ext == extra {
			return i + 1
		}
	}
	return -1
}

func part1() int {

	var arr [][]rune
	prevline := ""
	ans := 0

	for {

		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" && prevline == "" {
			break
		}

		if line == "" {

			mirror := 100 * horizontalMirror(arr, 0)
			if mirror == -100 {
				mirror = verticalMirror(arr, 0)
			}

			var temp [][]rune
			arr = temp
		}

		if line != "" {
			arr = append(arr, []rune(line))
		}
		prevline = line
	}
	return ans
}

func part2() int {

	var arr [][]rune
	prevline := ""
	ans := 0

	for {

		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" && prevline == "" {
			break
		}

		if line == "" {

			mirror := 100 * horizontalMirror(arr, 1)
			if mirror == -100 {
				mirror = verticalMirror(arr, 1)
			}
			ans += mirror

			var temp [][]rune
			arr = temp
		}

		if line != "" {
			arr = append(arr, []rune(line))
		}
		prevline = line
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
