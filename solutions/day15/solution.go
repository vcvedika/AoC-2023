package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	first  string
	second int
}

func hashCode(str string) int {

	res := 0
	for _, elem := range str {
		res += int(elem)
		res *= 17
		res %= 256
	}
	return res
}

func part1() int {

	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')
	line = strings.TrimSpace(line)

	ans := 0
	arr := strings.Split(line, ",")
	for _, elem := range arr {
		ans += hashCode(elem)
	}
	return ans
}

func part2() int {

	input := bufio.NewReader(os.Stdin)
	line, _ := input.ReadString('\n')
	line = strings.TrimSpace(line)

	ans := 0
	arr := strings.Split(line, ",")
	mp := make(map[int][]Pair)

	for _, elem := range arr {
		label := ""
		number := ""
		dash := false
		for i := 0; i < len(elem); i++ {
			if elem[i] == '-' {
				dash = true
				break
			}
			if elem[i] == '=' {
				for j := i + 1; j < len(elem); j++ {
					number += string(elem[j])
				}
				break
			}
			label = label + string(elem[i])
		}
		boxnumber := hashCode(label)
		if dash {
			for i, elem := range mp[boxnumber] {
				if elem.first == label {
					mp[boxnumber] = append(mp[boxnumber][:i], mp[boxnumber][i+1:]...)
					break
				}
			}
		} else {
			val, _ := strconv.Atoi(number)
			flag := false
			for i, elem := range mp[boxnumber] {
				if elem.first == label {
					mp[boxnumber][i].second = val
					flag = true
				}
			}
			if !flag {
				mp[boxnumber] = append(mp[boxnumber], Pair{label, val})
			}
		}
	}

	for i := 0; i < 256; i++ {
		for indx, elem := range mp[i] {
			ans += (i + 1) * (indx + 1) * (elem.second)
		}
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
