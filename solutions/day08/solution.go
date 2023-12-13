package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	first, second string
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a * b) / gcd(a, b)
}

func position(s string, ch rune) int {
	for i, elem := range s {
		if elem == ch {
			return i
		}
	}
	return -1
}

func part1() int {

	input := bufio.NewReader(os.Stdin)
	directions, _ := input.ReadString('\n')
	directions = strings.TrimSpace(directions)

	val := strings.Fields(directions)
	directions = strings.Join(val, "")
	graph := make(map[string]Pair)

	line, _ := input.ReadString('\n')
	line = strings.TrimSpace(line)

	for {
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		a := line[0:3]
		b := line[position(line, '(')+1 : position(line, '(')+4]
		c := line[position(line, ')')-3 : position(line, ')')]
		graph[a] = Pair{b, c}
	}

	i := 0
	ans := 0
	l := len(directions)
	temp := "AAA"

	for temp != "ZZZ" {
		if rune(directions[i]) == 'R' {
			temp = graph[temp].second
		} else {
			temp = graph[temp].first
		}
		i = (i + 1) % l
		ans = ans + 1
	}
	return ans
}

func part2() int {

	input := bufio.NewReader(os.Stdin)
	directions, _ := input.ReadString('\n')
	directions = strings.TrimSpace(directions)

	val := strings.Fields(directions)
	directions = strings.Join(val, "")
	graph := make(map[string]Pair)

	line, _ := input.ReadString('\n')
	line = strings.TrimSpace(line)

	var nodes []string

	for {
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		a := line[0:3]
		b := line[position(line, '(')+1 : position(line, '(')+4]
		c := line[position(line, ')')-3 : position(line, ')')]
		graph[a] = Pair{b, c}
		nodes = append(nodes, a)
	}

	i := 0
	ans := 1
	l := len(directions)

	for _, node := range nodes {
		if node[2] == 'A' {
			temp := node
			cnt := 0
			for temp[2] != 'Z' {
				if directions[i] == 'R' {
					temp = graph[temp].second
				} else {
					temp = graph[temp].first
				}
				i = (i + 1) % l
				cnt = cnt + 1
			}
			ans = lcm(ans, cnt)
		}
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
