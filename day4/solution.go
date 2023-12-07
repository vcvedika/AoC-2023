package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

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

		pos++
		var winning, present string
		var flag bool
		for i := pos; i < len(line); i++ {
			if line[i] == '|' {
				flag = true
				continue
			}

			if flag {
				present += string(line[i])
			}
			if !flag {
				winning += string(line[i])
			}
		}

		words := strings.Fields(winning)
		wordCounts := make(map[string]int)
		for _, word := range words {
			wordCounts[word]++
		}

		score := 1
		flag = true
		words = strings.Fields(present)
		for _, word := range words {
			if wordCounts[word] >= 1 {
				if flag {
					flag = false
					continue
				}
				score *= 2
			}
		}

		if score == 1 && flag {
			score = 0
		}

		ans += score
	}

	return ans
}

func part2() int {

	ans := 0
	mp := make(map[int]int)
	card := 0

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}

		card++
		mp[card]++
		pos := 5
		for unicode.IsDigit(rune(line[pos])) {
			pos++
		}

		pos++
		var winning, present string
		var flag bool
		for i := pos; i < len(line); i++ {
			if line[i] == '|' {
				flag = true
				continue
			}

			if flag {
				present += string(line[i])
			}
			if !flag {
				winning += string(line[i])
			}
		}

		words := strings.Fields(winning)
		wordCounts := make(map[string]int)
		for _, word := range words {
			wordCounts[word]++
		}

		score := 0
		words = strings.Fields(present)
		for _, word := range words {
			if wordCounts[word] >= 1 {
				score++
			}
		}

		for i := 0; i < score; i++ {
			mp[card+i+1] += mp[card]
		}
	}

	for key, val := range mp {
		if key <= card {
			ans += val
		}
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
