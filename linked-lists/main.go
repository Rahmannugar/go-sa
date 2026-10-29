package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

// addToHead inserts a value at the beginning of the list.
// Time complexity: O(1), Space complexity: O(1)
func (l *LinkedList[T]) addToHead(value T) {
	newNode := &Node[T]{value: value}

	newNode.next = l.head
	l.head = newNode
}

// addToTail inserts a value at the end of the list.
// Time complexity: O(n), Space complexity: O(1)
func (l *LinkedList[T]) addToTail(value T) {
	newNode := &Node[T]{value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// removeFromHead removes and returns the first value, or false if the list is empty.
// Time complexity: O(1), Space complexity: O(1)
func (l *LinkedList[T]) removeFromHead() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}

	value := l.head.value
	l.head = l.head.next

	return value, true
}

// forEach calls fn for every value from head to tail.
// Time complexity: O(n), Space complexity: O(1)
func (l *LinkedList[T]) forEach(fn func(T)) {
	current := l.head

	for current != nil {
		fn(current.value)
		current = current.next
	}
}

func main() {
	list := LinkedList[int]{}

	list.addToTail(10)
	list.addToTail(20)
	list.addToTail(30)
	list.addToTail(40)
	list.addToHead(33)

	removed, ok := list.removeFromHead()
	if ok {
		fmt.Println("Removed:", removed)
	}

	fmt.Print("Linked list: ")
	list.forEach(func(value int) {
		fmt.Printf("%d ", value)
	})
	fmt.Println()
}
