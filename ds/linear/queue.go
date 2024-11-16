package linear

import (
	"errors"
	"fmt"
)

// Queue represents a queue data structure.
type Queue[T comparable] struct {
	Front *Node[T]
	Back  *Node[T]
}

// IsEmpty checks if the stack is empty.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.Front == nil
}

// Enqueue adds a new element to the end of the Queue
func (queue *Queue[T]) Enqueue(data T) T {
	if queue.IsEmpty() {
		newNode := &Node[T]{data: data}
		queue.Back = newNode
		queue.Front = newNode
		return data
	}

	newBack := &Node[T]{data: data}
	queue.Back.next = newBack
	queue.Back = newBack

	return data
}

// Dequeue removes a new element to the beginning of the Queue
func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.IsEmpty() {
		var zero T
		return zero, errors.New("error: queue is empty")
	}

	queue.Front = queue.Front.next
	if queue.Front == nil {
		queue.Back = nil

		var zero T
		return zero, nil
	} else {
		return queue.Front.data, nil
	}
}

// Display prints all elements in the linked list.
func (queue *Queue[T]) Display() error {
	if queue.IsEmpty() {
		return errors.New("error: queue is empty")
	}

	current := queue.Front

	for current.next != nil {
		fmt.Printf("%v <- ", current.data)
		current = current.next
	}

	fmt.Printf("%v\n", current.data)

	return nil
}
