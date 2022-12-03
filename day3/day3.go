package main

import (
	utils "aoc2022/utils"
	"fmt"
	"sort"
	"strings"

	"github.com/meirf/gopart"
)

func main() {
	content := utils.ReadFromFile("input")
	inputs := strings.Split(string(content), "\n")

	res1 := part1(inputs); fmt.Println(res1)
	res2 := part2(inputs); fmt.Println(res2)
}

func part1(inputs []string) int {
	score := 0
	for _, val := range inputs {
		first := strings.Split(val[:len(val)/2], ""); sort.Strings(first)
		second := strings.Split(val[len(val)/2:], ""); sort.Strings(second)

		set := make(map[string]int)
		for i, j := 0, 0; i < len(first) && j < len(second); {
			switch {
			case first[i] == second[j]:
				if int(first[i][0]) >= 'a' {
					set[first[i]] = int(first[i][0]) - 'a' + 1
				} else {
					set[first[i]] = int(first[i][0]) - 'A' + 27
				}
				i++; j++
			case first[i] < second[j]:
				i++
			default:
				j++
			}
		}
		for _, val := range set {
			score += val
		}
	}
	return score
}

func part2(inputs []string) int {
	score := 0; var groups [][]string
	for i := range gopart.Partition(len(inputs), 3) {
		groups = append(groups, inputs[i.Low:i.High])
	}
	for _, group := range groups {
		res := utils.Map(group, func(s string) string {
			res := strings.Split(s, ""); sort.Strings(res)
			return strings.Join(res, "")
		})
		first, second, third := res[0], res[1], res[2]

		for i, j, k := 0, 0, 0; i < len(first) && j < len(second) && k < len(third); {
			if first[i] == second[j] && second[j] == third[k] {
				if int(first[i]) >= 'a' {
					score += int(first[i]) - 'a' + 1
				} else {
					score += int(first[i]) - 'A' + 27
				}
				break
			} else if first[i] < second[j] {
				i++
			} else if second[j] < third[k] {
				j++
			} else {
				k++
			}
		}
	}
	return score
}
