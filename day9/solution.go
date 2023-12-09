package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func allZero(arr []int) bool {
	for _, elem := range arr {
		if elem != 0 {
			return false
		}
	}
	return true
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

		data := strings.Fields(line)

		var diff, seq, last []int
		for _, elem := range data {
			val, _ := strconv.Atoi(elem)
			seq = append(seq, val)
		}

		last = append(last, seq[len(seq)-1])

		for allZero(seq) != true {
			for i := 1; i < len(seq); i++ {
				diff = append(diff, seq[i]-seq[i-1])
			}
			last = append(last, diff[len(diff)-1])
			seq = diff
			diff = []int{}
		}

		for _, elem := range last {
			ans += elem
		}
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

		data := strings.Fields(line)

		var diff, seq, first []int
		for _, elem := range data {
			val, _ := strconv.Atoi(elem)
			seq = append(seq, val)
		}

		first = append(first, seq[0])

		for allZero(seq) != true {
			for i := 1; i < len(seq); i++ {
				diff = append(diff, seq[i]-seq[i-1])
			}
			first = append(first, diff[0])
			seq = diff
			diff = []int{}
		}

		k := 1
		for _, elem := range first {
			ans += elem * k
			k = k * -1
		}
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
