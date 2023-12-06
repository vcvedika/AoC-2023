package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1() int {

	ans := 0

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		pos := 5
		for unicode.IsDigit(rune(line[pos])) {
			pos++
		}

		gameID, _ := strconv.Atoi(line[5:pos])
		var r, g, b int

		for i := pos; i < len(line); i++ {

			var initial int
			if unicode.IsDigit(rune(line[i])) {
				initial = i
				for unicode.IsDigit(rune(line[i])) {
					i++
				}
				i++
				if i+3 <= len(line) && line[i:i+3] == "red" {
					r, _ = strconv.Atoi(line[initial : i-1])
				}

				if i+5 <= len(line) && line[i:i+5] == "green" {
					g, _ = strconv.Atoi(line[initial : i-1])
				}

				if i+4 <= len(line) && line[i:i+4] == "blue" {
					b, _ = strconv.Atoi(line[initial : i-1])
				}
			}

			if r > 12 || g > 13 || b > 14 {
				gameID = 0
				break
			}
		}

		ans += gameID
	}

	return ans
}

func part2() int {

	ans := 0

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		pos := 5
		for unicode.IsDigit(rune(line[pos])) {
			pos++
		}

		var r, g, b, red, green, blue int

		for i := pos; i < len(line); i++ {

			var initial int
			if unicode.IsDigit(rune(line[i])) {
				initial = i
				for unicode.IsDigit(rune(line[i])) {
					i++
				}
				i++
				if i+3 <= len(line) && line[i:i+3] == "red" {
					r, _ = strconv.Atoi(line[initial : i-1])
				}

				if i+5 <= len(line) && line[i:i+5] == "green" {
					g, _ = strconv.Atoi(line[initial : i-1])
				}

				if i+4 <= len(line) && line[i:i+4] == "blue" {
					b, _ = strconv.Atoi(line[initial : i-1])
				}
			}
			red = max(red, r)
			green = max(green, g)
			blue = max(blue, b)
		}

		ans += red * green * blue
	}

	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
