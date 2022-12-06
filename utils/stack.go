package aoc2022

type Stack []int


func (s *Stack) Push(v int) {
	*s = append(*s, v) 
}


func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return -1, false
	} else {
		index := len(*s) - 1 
		res := (*s)[index] 
		*s = (*s)[:index] 
		return res, true
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) PopN(n int) ([]int, bool) {
	if len(*s) < n {
		return nil, false
	} else {
		index := len(*s) - n
		res := (*s)[index:]
		*s = (*s)[:index]
		return res, true
	}
}

func (s *Stack) PushN(n []int) {
	*s = append(*s, n...)
}