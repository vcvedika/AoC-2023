package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func mappedValue(mp map[int]Pair, a int) int {
	var res, flag int
	for key, value := range mp {
		if a >= key && a < key+value.second {
			res = value.first + a - key
			flag = 1
			break
		}
	}
	if flag == 0 {
		res = a
	}
	return res
}

func reverseMapping(mp map[int]Pair, a int) int {
	var res, flag int
	for key, value := range mp {
		if a >= value.first && a < value.first+value.second {
			res = key + a - value.first
			flag = 1
			break
		}
	}
	if flag == 0 {
		res = a
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func seedExists(seeds []string, seed int) bool {
	for i := 0; i < len(seeds); i += 2 {
		a, _ := strconv.Atoi(seeds[i])
		b, _ := strconv.Atoi(seeds[i+1])
		if seed >= a && seed < a+b {
			return true
		}
	}
	return false
}

type Pair struct {
	first, second int
}

func part1() int {

	ssmp := make(map[int]Pair)
	sfmp := make(map[int]Pair)
	fwmp := make(map[int]Pair)
	wlmp := make(map[int]Pair)
	ltmp := make(map[int]Pair)
	thmp := make(map[int]Pair)
	hlmp := make(map[int]Pair)

	var seeds []string
	var ss, sf, fw, wl, lt, th, hl bool

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if hl && line == "" {
			break
		}

		if len(line) > 0 && line[0:5] == "seeds" {
			seeds = strings.Fields(line[6:])
		}
		if len(line) > 0 && line[0:5] == "seed-" {
			ss = true
			continue
		}
		if len(line) > 0 && line[0:4] == "soil" {
			ss = false
			sf = true
			continue
		}
		if len(line) > 0 && line[0] == 'f' {
			sf = false
			fw = true
			continue
		}
		if len(line) > 0 && line[0] == 'w' {
			fw = false
			wl = true
			continue
		}
		if len(line) > 0 && line[0] == 'l' {
			wl = false
			lt = true
			continue
		}
		if len(line) > 0 && line[0] == 't' {
			lt = false
			th = true
			continue
		}
		if len(line) > 0 && line[0] == 'h' {
			th = false
			hl = true
			continue
		}

		data := strings.Fields(line)
		if len(data) == 0 {
			continue
		}
		source, _ := strconv.Atoi(data[0])
		dest, _ := strconv.Atoi(data[1])
		range_len, _ := strconv.Atoi(data[2])

		if ss {
			ssmp[dest] = Pair{source, range_len}
		}
		if sf {
			sfmp[dest] = Pair{source, range_len}
		}
		if fw {
			fwmp[dest] = Pair{source, range_len}
		}
		if wl {
			wlmp[dest] = Pair{source, range_len}
		}
		if lt {
			ltmp[dest] = Pair{source, range_len}
		}
		if th {
			thmp[dest] = Pair{source, range_len}
		}
		if hl {
			hlmp[dest] = Pair{source, range_len}
		}
	}

	location := math.MaxInt64
	for _, seed := range seeds {
		seed_value, _ := strconv.Atoi(seed)
		soil := mappedValue(ssmp, seed_value)
		fertilizer := mappedValue(sfmp, soil)
		water := mappedValue(fwmp, fertilizer)
		light := mappedValue(wlmp, water)
		temp := mappedValue(ltmp, light)
		humidity := mappedValue(thmp, temp)
		location_temp := mappedValue(hlmp, humidity)
		location = min(location, location_temp)
	}
	return location
}

//Part 2 needs to be optimized
func part2() int {

	ssmp := make(map[int]Pair)
	sfmp := make(map[int]Pair)
	fwmp := make(map[int]Pair)
	wlmp := make(map[int]Pair)
	ltmp := make(map[int]Pair)
	thmp := make(map[int]Pair)
	hlmp := make(map[int]Pair)

	var seeds []string
	var ss, sf, fw, wl, lt, th, hl bool

	for {
		input := bufio.NewReader(os.Stdin)
		line, _ := input.ReadString('\n')
		line = strings.TrimSpace(line)

		if hl && line == "" {
			break
		}

		if len(line) > 0 && line[0:5] == "seeds" {
			seeds = strings.Fields(line[6:])
		}
		if len(line) > 0 && line[0:5] == "seed-" {
			ss = true
			continue
		}
		if len(line) > 0 && line[0:4] == "soil" {
			ss = false
			sf = true
			continue
		}
		if len(line) > 0 && line[0] == 'f' {
			sf = false
			fw = true
			continue
		}
		if len(line) > 0 && line[0] == 'w' {
			fw = false
			wl = true
			continue
		}
		if len(line) > 0 && line[0] == 'l' {
			wl = false
			lt = true
			continue
		}
		if len(line) > 0 && line[0] == 't' {
			lt = false
			th = true
			continue
		}
		if len(line) > 0 && line[0] == 'h' {
			th = false
			hl = true
			continue
		}

		data := strings.Fields(line)
		if len(data) == 0 {
			continue
		}
		source, _ := strconv.Atoi(data[0])
		dest, _ := strconv.Atoi(data[1])
		range_len, _ := strconv.Atoi(data[2])

		if ss {
			ssmp[dest] = Pair{source, range_len}
		}
		if sf {
			sfmp[dest] = Pair{source, range_len}
		}
		if fw {
			fwmp[dest] = Pair{source, range_len}
		}
		if wl {
			wlmp[dest] = Pair{source, range_len}
		}
		if lt {
			ltmp[dest] = Pair{source, range_len}
		}
		if th {
			thmp[dest] = Pair{source, range_len}
		}
		if hl {
			hlmp[dest] = Pair{source, range_len}
		}
	}

	loc := 0

	for {
		humidity := reverseMapping(hlmp, loc)
		temp := reverseMapping(thmp, humidity)
		light := reverseMapping(ltmp, temp)
		water := reverseMapping(wlmp, light)
		fertilizer := reverseMapping(fwmp, water)
		soil := reverseMapping(sfmp, fertilizer)
		seed_value := reverseMapping(ssmp, soil)
		isPresent := seedExists(seeds, seed_value)

		if isPresent {
			break
		}
		loc++
	}
	return loc
}

func main() {
	fmt.Println(part1())
	// fmt.Println(part2())
}
