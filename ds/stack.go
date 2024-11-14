package ds

import "fmt"

// Stack represents a stack data structure with a pointer to the top node.
type Stack struct {
	top *Node
}

// IsEmpty checks if the stack is empty.
func (stack *Stack) IsEmpty() bool {
	return stack.top == nil
}

// Push adds a new element to the top of the stack.
func (stack *Stack) Push(data int) {
	if stack.IsEmpty() {
		stack.top = &Node{data: data}
		return
	}
	temp := stack.top
	stack.top = &Node{data: data}
	stack.top.next = temp
}

// Pop removes the top element from the stack.
func (stack *Stack) Pop() {
	if stack.IsEmpty() {
		return
	}

	temp := stack.top.next
	stack.top = temp
}

// Display prints all elements in the linked list.
func (stack *Stack) Display() {
	if stack.IsEmpty() {
		fmt.Println("Empty")
	}

	current := stack.top

	for current.next != nil {
		fmt.Printf("%d <- ", current.data)
		current = current.next
	}

	fmt.Printf("%d\n", current.data)
}
