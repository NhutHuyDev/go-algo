package main

import "github.com/NhutHuyDev/go-algo/ds"

func main() {
	/* Binary Search */
	// arr := []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}
	// fmt.Println(search.Binary_Search(arr, 32))

	/* Selection Sort */
	// arr := []int{75, 67, 42, 11, 13, 88, 7, 32, 56, 29, 12, 3, 66}
	// fmt.Println(sort.Selection_Sort(arr))

	/* Linked List */
	// lL := ds.LinkedList{}
	// lL.AddToBegining(20)
	// lL.AddToEnd(74)
	// lL.AddToEnd(33)
	// lL.AddToEnd(19)
	// lL.AddToEnd(21)
	// lL.AddToEnd(80)
	// lL.AddToEnd(52)
	// lL.AddToEnd(44)
	// lL.Display()

	/* Stack */
	stack := ds.Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Display()
}
