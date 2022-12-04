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

func Sum[T ~string | ~int](list []T) int {
	sum := 0
	for _, val := range list {
		switch any(val).(type) {
		case int:
			sum += any(val).(int)
		case string:
			num, _ := strconv.Atoi(any(val).(string))
			sum += num
		}
	}
	return sum
}

// stolen from https://gobyexample.com/collection-functions 
// tweaked to work with generics
func Map[T, V ~string | ~int](vs []T, f func(T) V) []V {
	vsm := make([]V, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
