package aoc2022

import (
	"io/ioutil"
	"log"
	"strconv"
)

func ReadFromFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func Sum[T any](list []T) int {
	sum := 0
	for _, val := range list {
		if v, ok := any(val).(int); ok {
			sum += v
		} else if v, ok := any(val).(string); ok {
			n, _ := strconv.Atoi(v)
			sum += n
		}
	}
	return sum
}
