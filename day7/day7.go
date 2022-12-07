package main

import (
	utils "aoc2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	Name	string
	Size	int
	Children []*Node
	Parent 	*Node
}
func main() {
	content := utils.ReadFromFile("input.txt")
	inputs := strings.Split(string(content), "\n")

	curr := &Node{Name: "root", Parent: nil, Size : 0}
	root := curr
	for i := 1; i < len(inputs); i++ {
		if strings.Contains(inputs[i], "$ cd") {
			name := inputs[i][5:]
			fmt.Println(name)
			if name == ".." {
				curr = curr.Parent
			} else {
				par := curr
				curr = &Node{Name: inputs[i][5:], Parent: par}
				par.Children = append(par.Children, curr)
			}
		} else if strings.Contains(inputs[i], "$ ls") {
			for j := i+1; j < len(inputs); j++ {
				if strings.Contains(inputs[j], "$") {
					i = j - 1
					break
				}
				if strings.Contains(inputs[j], "dir") {
				} else {
					temp, _ := strconv.Atoi(strings.Split(inputs[j], " ")[0])
					curr.Children = append(curr.Children, &Node{Name: strings.Split(inputs[j], " ")[1], Parent: curr, Size: temp})
					//curr.Size += temp
				}
			}
		}
	}
	total, res := 0, 0
	part1(root, &total)
	fmt.Println("root size: ", root.Size)
	fmt.Println("part1: ", total)

	needed := 30000000 - math.Abs(float64(70000000 - root.Size))
	part2(root, int(needed), &res)
	fmt.Println("part2: ", res)
}

func part1(root *Node, total *int) {
	for _, node := range root.Children {
		part1(node, total)
		root.Size += node.Size
	}
	if root.Size <= 100000 && root.Children != nil {
		*total += root.Size
	}
}

func part2(root *Node, needed int, res *int) {
	for _, node := range root.Children {
		part2(node, needed, res)
		if node.Size >= needed && node.Children != nil {
			if *res == 0 {
				*res = node.Size
			} else if node.Size < *res && *res != 0{
				*res = node.Size
			}
		}
	}
}