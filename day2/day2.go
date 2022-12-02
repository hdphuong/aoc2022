package main

import (
	utils "aoc2022/utils"
	"fmt"
	"strings"
)

func main() {
	content := utils.ReadFromFile("input")
	inputs := strings.Split(string(content), "\n")

	chanShape1, chanOutcome1 := make(chan int), make(chan int)
	chanShape2, chanOutcome2 := make(chan int), make(chan int)

	go shapeScore(inputs, chanShape1, chanShape2)
	go outcomeScore(inputs, chanOutcome1, chanOutcome2)

	fmt.Println("part 1 total: ", (<-chanShape1)+(<-chanOutcome1))
	fmt.Println("part 2 total: ", (<-chanShape2)+(<-chanOutcome2))
}

func shapeScore(inputs []string, chanShape1 chan int, chanShape2 chan int) {
	score1, score2 := 0, 0
	for _, input := range inputs {
		score1 += int(input[2]-'X') + 1

		// uhh things i wasted time on:
		// 1. type casting: 'X' - 'Y' would return 255 instead of -1 if I do int('X' - 'Y') since it's a byte.
		// 2. modulo: -2 % 3 == 0 in Go (wtf) not 1 as in math so needs to add 3 before doing modulo math
		o, u := int(int8(input[0]-'A')), int(int8(input[2]-'Y'))
		score2 += (o+(u+3))%3 + 1

	}
	chanShape1 <- score1
	chanShape2 <- score2

}

func outcomeScore(inputs []string, chanScore1 chan int, chanScore2 chan int) {
	score1, score2 := 0, 0
	for _, input := range inputs {
		o, u := input[0]-'A', input[2]-'X'
		score1 += (int(u-o+3+1) % 3) * 3 // plus 3 to make -2 % 3 == 1 and plus 1 cos we give points when tie
		score2 += int(input[2]-'X') * 3
	}
	chanScore1 <- score1
	chanScore2 <- score2
}
