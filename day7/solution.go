package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {

	rank := make(map[string]int)
	card := make(map[string]int)

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		data := strings.Fields(line)
		card[data[0]], _ = strconv.Atoi(data[1])

		cardno := data[0]
		mp := make(map[rune]int)
		for _, ch := range cardno {
			mp[ch]++
		}

		if len(mp) == 1 {
			rank[cardno] = 7
		} else if len(mp) == 2 {
			flag := false
			for _, val := range mp {
				if val == 1 || val == 4 {
					continue
				} else {
					flag = true
					break
				}
			}
			if !flag {
				rank[cardno] = 6
			} else {
				rank[cardno] = 5
			}
		} else if len(mp) == 3 {
			flag := false
			for _, val := range mp {
				if val == 1 || val == 3 {
					continue
				} else {
					flag = true
					break
				}
			}
			if !flag {
				rank[cardno] = 4
			} else {
				rank[cardno] = 3
			}
		} else if len(mp) == 4 {
			rank[cardno] = 2
		} else {
			rank[cardno] = 1
		}
	}

	list := make(map[int][]string)
	for key, value := range rank {
		list[value] = append(list[value], key)
	}

	ordering := make(map[rune]int)
	ordering['A'] = 13
	ordering['K'] = 12
	ordering['Q'] = 11
	ordering['J'] = 10
	ordering['T'] = 9
	ordering['9'] = 8
	ordering['8'] = 7
	ordering['7'] = 6
	ordering['6'] = 5
	ordering['5'] = 4
	ordering['4'] = 3
	ordering['3'] = 2
	ordering['2'] = 1

	for _, val := range list {
		for i := 0; i < len(val)-1; i++ {
			min_indx := i
			for j := i + 1; j < len(val); j++ {
				for k := 0; k < 5; k++ {
					if rune(val[j][k]) != rune(val[min_indx][k]) {
						if ordering[rune(val[j][k])] < ordering[rune(val[min_indx][k])] {
							min_indx = j
						}
						break
					}
				}
			}
			if min_indx != i {
				val[i], val[min_indx] = val[min_indx], val[i]
			}
		}
	}

	ans := 0
	k := 1

	for i := 1; i <= 7; i++ {
		for _, elem := range list[i] {
			ans += k * card[elem]
			k++
		}
	}
	return ans
}

//Unable to debug part 2
func part2() int {

	rank := make(map[string]int)
	card := make(map[string]int)

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		data := strings.Fields(line)
		card[data[0]], _ = strconv.Atoi(data[1])

		cardno := data[0]
		mp := make(map[rune]int)
		for _, ch := range cardno {
			mp[ch]++
		}

		if len(mp) == 1 {
			rank[cardno] = 7
		} else if len(mp) == 2 {
			flag := false
			for _, val := range mp {
				if val == 1 || val == 4 {
					continue
				} else {
					flag = true
					break
				}
			}
			if !flag {
				if mp['J'] == 1 || mp['J'] == 4 {
					rank[cardno] = 7
				} else {
					rank[cardno] = 6
				}
			} else {
				if mp['J'] == '2' || mp['J'] == '3' {
					rank[cardno] = 7
				} else {
					rank[cardno] = 5
				}
			}
		} else if len(mp) == 3 {
			flag := false
			for _, val := range mp {
				if val == 1 || val == 3 {
					continue
				} else {
					flag = true
					break
				}
			}
			if !flag {
				if mp['J'] == 1 || mp['J'] == 3 {
					rank[cardno] = 6
				} else {
					rank[cardno] = 4
				}
			} else {
				if mp['J'] == 2 {
					rank[cardno] = 6
				} else if mp['J'] == 1 {
					rank[cardno] = 5
				} else {
					rank[cardno] = 3
				}
			}
		} else if len(mp) == 4 {
			if mp['J'] == 2 || mp['J'] == 1 {
				rank[cardno] = 4
			} else {
				rank[cardno] = 2
			}
		} else {
			if mp['J'] == 1 {
				rank[cardno] = 2
			} else {
				rank[cardno] = 1
			}
		}
	}

	list := make(map[int][]string)
	for key, value := range rank {
		list[value] = append(list[value], key)
	}

	ordering := make(map[rune]int)
	ordering['A'] = 13
	ordering['K'] = 12
	ordering['Q'] = 11
	ordering['T'] = 10
	ordering['9'] = 9
	ordering['8'] = 8
	ordering['7'] = 7
	ordering['6'] = 6
	ordering['5'] = 5
	ordering['4'] = 4
	ordering['3'] = 3
	ordering['2'] = 2
	ordering['J'] = 1

	for _, val := range list {
		for i := 0; i < len(val)-1; i++ {
			min_indx := i
			for j := i + 1; j < len(val); j++ {
				for k := 0; k < 5; k++ {
					if rune(val[j][k]) != rune(val[min_indx][k]) {
						if ordering[rune(val[j][k])] < ordering[rune(val[min_indx][k])] {
							min_indx = j
						}
						break
					}
				}
			}
			if min_indx != i {
				val[i], val[min_indx] = val[min_indx], val[i]
			}
		}
	}

	ans := 0
	k := 1

	for i := 1; i <= 7; i++ {
		for _, elem := range list[i] {
			ans += k * card[elem]
			k++
		}
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
