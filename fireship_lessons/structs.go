package main

import "fmt"

// Define a stack type using a struct
type stack struct {
	index int
	data  [5]int
}

// Define the push and pop methods
func (s *stack) push(k int) {
	s.data[s.index] = k
	s.index++
}

// Here the stack pointer s is passed as an argument
func (s *stack) pop() int {
	s.index--
	return s.data[s.index]
}

func main() {
	// Create a pointer to the new stack and push 2 values
	s := new(stack)
	s.push(23)
	s.push(14)
	fmt.Println("Stack: \n", *s)
}
