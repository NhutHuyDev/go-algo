package ds

import "fmt"

// Node represents a single node in the linked list.
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list with a head pointer.
type LinkedList struct {
	head *Node
}

func (lL *LinkedList) AddToBegining(data int) {
	newNode := &Node{data: data}
	lL.head = newNode
}

// AddToEnd adds a new node with the specified value to the end of the linked list.
func (lL *LinkedList) AddToEnd(data int) {
	newNode := &Node{data: data}

	if lL.head == nil {
		lL.head = newNode
	} else {
		current := lL.head
		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

// Remove removes the first node with the specified value from the linked list.
func (lL *LinkedList) Remove(data int) {
	if lL.head == nil {
		return
	}

	if lL.head.data == data {
		lL.head = lL.head.next
		return
	}

	current := lL.head
	for current.next != nil {

		if current.next.data == data {
			current.next = current.next.next
			return
		}

		current = current.next
	}
}

// Display prints all elements in the linked list.
func (lL *LinkedList) Display() {
	current := lL.head
	for current.next != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Printf("%d -> nil \n", current.data)
}
