package main

import "fmt"

func NorthBeams(arr [][]rune) int {

	ans := 0
	row := len(arr)
	col := len(arr[0])
	mp := make(map[int]int)

	for i := 0; i < col; i++ {
		var seq []int
		var cnt []int
		str := "#"
		temp := 0
		for j := 0; j < row; j++ {
			str += string(rune(arr[j][i]))
		}
		str += "#"
		for in, elem := range str {
			if rune(elem) == 'O' {
				temp++
			} else if rune(elem) == '#' {
				seq = append(seq, in)
				cnt = append(cnt, temp)
				temp = 0
			}
		}
		for j := 1; j < len(seq); j++ {
			for k := 0; k < cnt[j]; k++ {
				mp[seq[j-1]+k]++
			}
		}
	}

	for i := 0; i < row; i++ {
		ans += mp[i] * col
		col--
	}

	return ans
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

	return NorthBeams(arr)
}

func main() {
	fmt.Println(part1())
}
