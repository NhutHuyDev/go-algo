package linear

import (
	"errors"
	"fmt"
)

// LinkedList represents the linked list with a head pointer.
type LinkedList[T comparable] struct {
	head *Node[T]
}

// AddToEnd adds a new node with the specified value to the head of the linked list.
func (lL *LinkedList[T]) AddToBegining(data T) {
	newNode := &Node[T]{data: data}
	lL.head = newNode
}

// AddToEnd adds a new node with the specified value to the end of the linked list.
func (lL *LinkedList[T]) AddToEnd(data T) T {
	newNode := &Node[T]{data: data}

	if lL.head == nil {
		lL.AddToBegining(data)
	} else {
		current := lL.head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}

	return data
}

// Remove removes the first node with the specified value from the linked list.
func (lL *LinkedList[T]) Remove(data T) (T, error) {
	if lL.head == nil {
		var zero T
		return zero, errors.New("error: list is empty")
	}

	if lL.head.data == data {
		lL.head = lL.head.next
		return data, nil
	}

	current := lL.head
	for current.next != nil {

		if current.next.data == data {
			current.next = current.next.next
			return data, nil
		}

		current = current.next
	}

	var zero T
	return zero, errors.New("error: data value is not found")
}

// Display prints all elements in the linked list.
func (lL *LinkedList[T]) Display() error {
	if lL.head == nil {
		return errors.New("error: list is empty")
	}

	current := lL.head
	for current.next != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}

	fmt.Printf("%v -> nil \n", current.data)
	return nil
}
