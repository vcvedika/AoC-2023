package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() int {

	ans := 1

	input := bufio.NewReader(os.Stdin)
	time, _ := input.ReadString('\n')
	time = strings.TrimSpace(time)

	input = bufio.NewReader(os.Stdin)
	distance, _ := input.ReadString('\n')
	distance = strings.TrimSpace(distance)

	time_data := strings.Fields(time[6:])
	distance_data := strings.Fields(distance[10:])

	for i, val := range time_data {
		t, _ := strconv.Atoi(val)
		d, _ := strconv.Atoi(distance_data[i])

		cnt := 0
		for j := 1; j <= t; j++ {
			if ((t - j) * j) > d {
				cnt++
			}
		}

		ans *= cnt
	}
	return ans
}

func part2() int {

	ans := 0

	input := bufio.NewReader(os.Stdin)
	time, _ := input.ReadString('\n')
	time = strings.TrimSpace(time)

	input = bufio.NewReader(os.Stdin)
	distance, _ := input.ReadString('\n')
	distance = strings.TrimSpace(distance)

	time_data := strings.Fields(time[6:])
	distance_data := strings.Fields(distance[10:])
	time = strings.Join(time_data, "")
	distance = strings.Join(distance_data, "")

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)

	for i := 1; i <= t; i++ {
		if ((t - i) * i) > d {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
