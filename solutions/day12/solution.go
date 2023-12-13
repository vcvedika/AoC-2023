package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func solve(i int, g int, pattern string, sequence []int, dp [][]int) int {

	if g == len(sequence) {
		if !strings.ContainsAny(pattern[i:], "#") {
			return 1
		}
		return 0
	}

	if i == len(pattern) {
		return 0
	}

	if dp[i][g] != -1 {
		return dp[i][g]
	}

	ans := 0

	if i+sequence[g] < len(pattern) {
		if !strings.ContainsAny(pattern[i:i+sequence[g]], ".") && pattern[i+sequence[g]] != '#' {
			ans += solve(i+sequence[g]+1, g+1, pattern, sequence, dp)
		}
	}

	if pattern[i] != '#' {
		ans += solve(i+1, g, pattern, sequence, dp)
	}

	dp[i][g] = ans
	return ans
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

		pattern := ""
		var sequence []int
		line += ","
		pos := 0

		for i, elem := range line {
			if unicode.IsDigit(rune(elem)) {
				pos = i
				break
			} else {
				if elem != ' ' {
					pattern += string(elem)
				}
			}
		}

		num := ""

		for i := pos; i < len(line); i++ {
			if unicode.IsDigit((rune(line[i]))) {
				num += string(line[i])
			} else {
				val, _ := strconv.Atoi(num)
				sequence = append(sequence, val)
				num = ""
			}
		}

		pattern += "."
		dp := make([][]int, len(pattern))
		for i := range dp {
			dp[i] = make([]int, len(sequence))
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		ans += solve(0, 0, pattern, sequence, dp)
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

		pattern := ""
		var sequence []int
		line += ","
		pos := 0

		for i, elem := range line {
			if unicode.IsDigit(rune(elem)) {
				pos = i
				break
			} else {
				if elem != ' ' {
					pattern += string(elem)
				}
			}
		}

		temp := pattern
		for i := 0; i < 4; i++ {
			pattern = pattern + "?" + temp
		}

		num := ""

		for i := pos; i < len(line); i++ {
			if unicode.IsDigit((rune(line[i]))) {
				num += string(line[i])
			} else {
				val, _ := strconv.Atoi(num)
				sequence = append(sequence, val)
				num = ""
			}
		}

		val := sequence
		for i := 0; i < 4; i++ {
			for _, elem := range val {
				sequence = append(sequence, elem)
			}
		}

		pattern += "."
		dp := make([][]int, len(pattern))
		for i := range dp {
			dp[i] = make([]int, len(sequence))
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}
		ans += solve(0, 0, pattern, sequence, dp)
	}

	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
