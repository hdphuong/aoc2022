package main

import (
	utils "aoc2022/utils"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	inputs := strings.Split(string(content), "\n")
	/*
	state := map[int]*utils.Stack{
		1: &utils.Stack{'Z', 'N'},
		2: &utils.Stack{'M', 'C', 'D'},
		3: &utils.Stack{'P'},
	} */
	state := map[int]*utils.Stack{
		1: {'F', 'C', 'J', 'P', 'H', 'T', 'W'},
		2: {'G', 'R', 'V', 'F', 'Z', 'J', 'B', 'H'},
		3: {'H', 'P', 'T', 'R'},
		4: {'Z', 'S', 'N', 'P', 'H', 'T'},
		5: {'N', 'V', 'F', 'Z', 'H', 'J', 'C', 'D'},
		6: {'P', 'M', 'G', 'F', 'W', 'D', 'Z'},
		7: {'M', 'V', 'Z', 'W', 'S', 'J', 'D', 'P'},
		8: {'N', 'D', 'S'},
		9: {'D', 'Z', 'S', 'F', 'M'},
	}
	//fmt.Println("part 1: ", part1(inputs, state))
	fmt.Println("part 2: ", part2(inputs, state))

}

func part1(inputs []string, state map[int]*utils.Stack) string {
	for _, input := range inputs {
		res := SplitSplit(input)
		for i := 0; i < res[0]; i++ {
			if v, ok := state[res[1]].Pop(); ok{
				state[res[2]].Push(v)
			}
		}
	}
	var ret string
	for z := 1; z < 4; z++{
		if !state[z].IsEmpty() {
			v, _ := state[z].Pop()
			ret += string(v)
		}
	}
	return ret
}

func part2(inputs []string, state map[int]*utils.Stack) string {
	var ret string
	for _, input := range inputs {
		res := SplitSplit(input)
		fmt.Println(state[1], state[2], state[3])
		if v, ok := state[res[1]].PopN(res[0]); ok{
			state[res[2]].PushN(v)
		}
	}
	for z := 1; z < 10; z++{
		if !state[z].IsEmpty() {
			v, _ := state[z].Pop()
			ret += string(v)
		}
	}
	return ret

}

func SplitSplit(input string) []int {
	temp := strings.Split(input, " ")
	res := utils.Map(temp, func(x string) int {
		num, err := strconv.Atoi(x)
		if err == nil {
			return num 
		} else {
			return -1
		}
	})
	return []int{res[1], res[3], res[5]}
}




