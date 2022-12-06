package main

import (
	utils "aoc2022/utils"
	"fmt"
	"strings"
)
func main() {
	content := utils.ReadFromFile("input.txt")
	fmt.Println("part 1: ", partX(strings.Split(string(content), "\n")[0], 4))
	fmt.Println("part 2: ", partX(strings.Split(string(content), "\n")[0], 14))
}

func partX(input string, num int) int {
	i := 0;
	var list []byte
	for i < len(input) {
		enqueue(&list, input[i], num)
		check := true
		for j := 0; j < len(list); j++ {
			for k := j + 1; k < len(list); k++ {
				if list[j] == list[k] {
					check = false
				}
			}
		}
		i++
		if check && len(list) == num {
			fmt.Println(list)
			return i 
		}

	}
	return i
}

func enqueue(q *[]byte, v byte, num int) {
	if len(*q) < num {
		*q = append(*q, v)
	} else {
		*q = append((*q)[1:], v)
	}
}
