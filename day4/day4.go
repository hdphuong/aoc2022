package main

import (
	utils "aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	content := utils.ReadFromFile("input")
	inputs := strings.Split(string(content), "\n")

	fmt.Println("part 1: ", part1(inputs))
	fmt.Println("part 2: ", part2(inputs))

}

func part1(inputs []string) int {
	count := 0 
	for _, input := range inputs {
		res := SplitSplit(input)	
		if res[0] <= res[2] && res[1] >= res[3] {
			count++
		} else if res[0] >= res[2] && res[1] <= res[3] {
			count++
		}
		
	}	
	return count
}


func part2(inputs []string) int {
	count := 0
	for _, input := range inputs {
		arr := [100]bool{false}
		res := SplitSplit(input)

		for i := res[0]; i <= res[1]; i++ { arr[i] = true }
		for i := res[2]; i <= res[3]; i++ {
			if arr[i] { count++; break }
		}
	}
	return count
}


func SplitSplit(input string) []int {
	temp := strings.FieldsFunc(input, func(r rune) bool {
		return r == '-' || r == ','
	})
	res := utils.Map(temp, func(x string) int {
		num, _ := strconv.Atoi(x)
		return num
	})
	return res
}