package main

import "fmt"

type pair struct {
	first  int
	second int
}

type pairu64 struct {
	first  uint64
	second uint64
}

type stack struct {
    data []int
    pos int
}

var errEmpty = fmt.Errorf("stack empty")

func (s *stack) pop() (int, error){
    lastpos := len(s.data) -1
    if lastpos == -1 {
        return 0, errEmpty
    }

    result := s.data[lastpos]
    s.data = s.data[:lastpos]

    return result, nil
}

func (s *stack) push(data int) {
    s.data = append(s.data, data)
}

func (s *stack) size() int {
    return len(s.data)
}

func (s *stack) peek() int {
    return s.data[len(s.data)-1]
}

