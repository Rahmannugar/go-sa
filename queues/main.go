package main

import "fmt"

type Queue[T any] struct {
	items []T
}

// enqueue adds an item to the back of the queue.
// Time complexity: O(1) amortized, Space complexity: O(1) extra
func (q *Queue[T]) enqueue(item T) {
	q.items = append(q.items, item)
}

// dequeue removes and returns the front item, or false if the queue is empty.
// Time complexity: O(1), Space complexity: O(1)
func (q *Queue[T]) dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	val := q.items[0]
	q.items = q.items[1:]

	return val, true
}

// peek returns the front item without removing it, or false if the queue is empty.
// Time complexity: O(1), Space complexity: O(1)
func (q *Queue[T]) peek() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}

	return q.items[0], true
}

func main() {
	var queue Queue[int]
	queue.enqueue(10)
	queue.enqueue(20)
	queue.enqueue(33)

	fmt.Println(queue.peek())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.dequeue())
	fmt.Println(queue.dequeue())
}
