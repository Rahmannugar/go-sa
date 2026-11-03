package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type DoublyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// addToHead inserts a value at the beginning of the list.
// Time complexity: O(1), Space complexity: O(1)
func (l *DoublyLinkedList[T]) addToHead(value T) {
	newNode := &Node[T]{value: value}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}

	l.size++
}

// addToTail inserts a value at the end of the list.
// Time complexity: O(1), Space complexity: O(1)
func (l *DoublyLinkedList[T]) addToTail(value T) {
	newNode := &Node[T]{value: value}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	}

	l.size++
}

// removeFromHead removes and returns the first value, or false if the list is empty.
// Time complexity: O(1), Space complexity: O(1)
func (l *DoublyLinkedList[T]) removeFromHead() (T, bool) {
	if l.head == nil {
		var zero T
		return zero, false
	}

	value := l.head.value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.size--
	return value, true
}

// removeFromTail removes and returns the last value, or false if the list is empty.
// Time complexity: O(1), Space complexity: O(1)
func (l *DoublyLinkedList[T]) removeFromTail() (T, bool) {
	if l.tail == nil {
		var zero T
		return zero, false
	}

	value := l.tail.value
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	l.size--
	return value, true
}

// forEach calls fn for every value from head to tail.
// Time complexity: O(n), Space complexity: O(1)
func (l *DoublyLinkedList[T]) forEach(fn func(T)) {
	current := l.head

	for current != nil {
		fn(current.value)
		current = current.next
	}
}

// forEachReverse calls fn for every value from tail to head.
// Time complexity: O(n), Space complexity: O(1)
func (l *DoublyLinkedList[T]) forEachReverse(fn func(T)) {
	current := l.tail

	for current != nil {
		fn(current.value)
		current = current.prev
	}
}

// sizeOf returns the number of nodes.
// Time complexity: O(1), Space complexity: O(1)
func (l *DoublyLinkedList[T]) sizeOf() int {
	return l.size
}

func main() {
	list := DoublyLinkedList[int]{}

	list.addToTail(10)
	list.addToTail(20)
	list.addToTail(30)
	list.addToHead(5)

	fmt.Print("Forward: ")
	list.forEach(func(value int) {
		fmt.Printf("%d ", value)
	})
	fmt.Println()

	fmt.Print("Backward: ")
	list.forEachReverse(func(value int) {
		fmt.Printf("%d ", value)
	})
	fmt.Println()

	removed, ok := list.removeFromHead()
	if ok {
		fmt.Println("Removed from head:", removed)
	}

	removed, ok = list.removeFromTail()
	if ok {
		fmt.Println("Removed from tail:", removed)
	}

	fmt.Print("Remaining: ")
	list.forEach(func(value int) {
		fmt.Printf("%d ", value)
	})
	fmt.Println()

	fmt.Println("Size:", list.sizeOf())
}
