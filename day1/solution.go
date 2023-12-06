package main

import (
	"fmt"
	"unicode"
)

func Reverse(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		result = string(s[i]) + result
	}
	return result
}

func isPresent(s string, words [10]string) int {
	for i, str := range words {
		if str == s {
			return i
		}
	}
	return -1
}

func solve(s string, words [10]string) int {

	digit := -1
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			digit = int(s[i] - '0')
			break
		}
		for j := 1; j <= 5; j++ {
			if i+j >= len(s) {
				break
			}
			substr := s[i : i+j]
			digit = isPresent(substr, words)
			if digit != -1 {
				break
			}
		}
		if digit != -1 {
			break
		}
	}

	return digit
}

func part1() int {

	sum := 0

	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}

		var first, last int
		flag := true

		for i := 0; i < len(input); i++ {
			if unicode.IsDigit(rune(input[i])) {
				if flag {
					first = int(input[i] - '0')
					flag = false
				}
				last = int(input[i] - '0')
			}
		}
		sum += first*10 + last
	}
	return sum
}

func part2() int {

	sum := 0
	digits := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for {
		var input string
		fmt.Scanln(&input)
		if input == "" {
			break
		}

		first := solve(input, digits)

		input = Reverse(input)
		for i := 0; i < 10; i++ {
			digits[i] = Reverse(digits[i])
		}

		last := solve(input, digits)
		sum += first*10 + last

		for i := 0; i < 10; i++ {
			digits[i] = Reverse(digits[i])
		}
	}
	return sum
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
