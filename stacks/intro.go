package main

type Stack[T any] struct {
	items []T
}

// push adds an item to the top of the stack.
// Time complexity: O(1) amortized, Space complexity: O(1)
func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

// peek returns the top item without removing it, or false if the stack is empty.
// Time complexity: O(1), Space complexity: O(1)
func (s *Stack[T]) peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

// pop removes and returns the top item, or false if the stack is empty.
// Time complexity: O(1), Space complexity: O(1)
func (s *Stack[T]) pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}

	lastIndex := len(s.items) - 1
	val := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return val, true
}

// size returns the number of items in the stack.
// Time complexity: O(1), Space complexity: O(1)
func (s *Stack[T]) size() int {
	return len(s.items)
}
