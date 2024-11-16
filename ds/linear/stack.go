package linear

import (
	"errors"
	"fmt"
)

// import "fmt"

// Stack represents a stack data structure with a pointer to the top node.
type Stack[T comparable] struct {
	top *Node[T]
}

// IsEmpty checks if the stack is empty.
func (stack *Stack[T]) IsEmpty() bool {
	return stack.top == nil
}

// Push adds a new element to the top of the stack.
func (stack *Stack[T]) Push(data T) T {
	if stack.IsEmpty() {
		stack.top = &Node[T]{data: data}
	} else {
		temp := stack.top
		stack.top = &Node[T]{data: data}
		stack.top.next = temp
	}

	return data
}

// Pop removes the top element from the stack.
func (stack *Stack[T]) Pop() (T, error) {
	if stack.IsEmpty() {
		var zero T
		return zero, errors.New("error: stack is empty")
	}

	stack.top = stack.top.next

	if stack.top == nil {
		var zero T
		return zero, nil
	} else {
		return stack.top.data, nil
	}
}

// Display prints all elements in the linked list.
func (stack *Stack[T]) Display() error {
	if stack.IsEmpty() {
		return errors.New("error: stack is empty")
	}

	current := stack.top

	for current.next != nil {
		fmt.Printf("%v <- ", current.data)
		current = current.next
	}

	fmt.Printf("%v\n", current.data)

	return nil
}
