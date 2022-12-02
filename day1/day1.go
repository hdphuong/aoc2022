package main

import (
	utils "aoc2022/utils"
	"fmt"
	"sort"
	"strings"
)

func main() {
	content := utils.ReadFromFile("input")
	res := strings.Split(string(content), "\n\n")

	var cals []int
	for _, elem := range res {
		nums := strings.Split(string(elem), "\n")
		cals = append(cals, utils.Sum(nums))
	}

	sort.Ints(cals)
	fmt.Println("sum: ", utils.Sum(cals[len(cals)-3:]))
}
