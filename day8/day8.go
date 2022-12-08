package main

import (
	utils "aoc2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	content := utils.ReadFromFile("input.txt")
	inputs := strings.Split(string(content), "\n")

	var forest [][]int
	for line := range inputs {
		list := utils.Map(strings.Split(inputs[line], ""), func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		forest = append(forest, list)
	}
	//fmt.Println(forest)

	result := make([][]bool, len(forest))
	for i := range result {
		result[i] = make([]bool, len(forest[i]))
	}
	part1(forest, result)
	part2(forest, result)
}

func part1(forest [][]int, result [][]bool) {
	for r := 0; r < len(forest); r++ {
		currMax1, currMax2 := -1, -1
		for c := 0; c < len(forest[r]); c++ {
			if forest[r][c] > currMax1 {
				currMax1 = forest[r][c]
				result[r][c] = true
			}
		}
		for c := len(forest[r]) - 1; c >= 0; c-- {
			if forest[r][c] > currMax2 {
				currMax2 = forest[r][c]
				result[r][c] = true
			}
		}
	}
	for c := 0; c < len(forest[0]); c++ {
		currMax1, currMax2 := -1, -1
		for r := 0; r < len(forest); r++ {
			if forest[r][c] > currMax1  {
				currMax1 = forest[r][c]
				result[r][c] = true
			}
		}
		for r := len(forest) - 1; r >= 0; r-- {
			if forest[r][c] > currMax2 {
				currMax2 = forest[r][c]
				result[r][c] = true
			}
		}
	}
	count := 0
	for _, v := range result {
		for _, z := range v {
			if z {
				count++
			}
		}
	}

	//fmt.Println(result)
	fmt.Println("part 1: ", count)
}

func part2 (forest [][]int, result [][]bool) {
	ret := 0
	for r := 0; r < len(result); r++ {
		for c := 0; c < len(result[r]); c++ {
			if !result[r][c] {
				continue
			}
			left, up := c, r 
			right := len(result[0]) - c - 1
			down := len(result) - r - 1
			for i := r - 1; i >= 0; i-- {
				if forest[i][c] >= forest[r][c] {
					up = r - i
					break
				}
			}
			for i := r + 1; i < len(result); i++ {
				if forest[i][c] >= forest[r][c] {
					down = i - r
					break
				}
			}
			for i := c - 1; i >= 0; i-- {
				if forest[r][i] >= forest[r][c] {
					left = c - i
					break
				}
			}
			for i := c + 1; i < len(result[r]); i++ {
				if forest[r][i] >= forest[r][c] {
					right = i - c
					break
				}
			}
			dist := left * right * up * down
			if dist > ret {
				ret = dist
			}
		}
	}
	fmt.Println("part 2: ", ret)
}
