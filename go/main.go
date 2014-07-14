package main

import (
	"fmt"
	//"strconv"
)

//simple stack structure
type stack []int

func (s stack) Empty() bool { return len(s) == 0 }
func (s stack) Peek() int   { return s[len(s)-1] }
func (s *stack) Put(i int)  { (*s) = append((*s), i) }
func (s *stack) Pop() int {
	d := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return d
}

//reverse polish notation calculator
//takes a string like "699+/"
//prints result e.g. 9
func ReversePolishCalculator(expression string) {
	var s stack
	for _, c := range expression {
		switch c {
		case '+':
			a := s.Pop()
			b := s.Pop()
			s.Put(a + b)
			break
		case '-':
			a := s.Pop()
			b := s.Pop()
			s.Put(a - b)
			break
		case '/':
			a := s.Pop()
			b := s.Pop()
			s.Put(a / b)
			break
		case '*':
			a := s.Pop()
			b := s.Pop()
			s.Put(a * b)
			break
		default:
			s.Put(int(c) - 48)
		}
	}

	//result
	fmt.Printf("%v=%v", expression, s.Pop())

}

func main() {
	ReversePolishCalculator("699+/")
}
