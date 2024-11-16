package main

import (
	"fmt"

	"github.com/NhutHuyDev/go-algo/ds/linear"
)

func main() {
	var err error

	/* Binary Search */
	// arr := []int{3, 7, 11, 12, 13, 29, 32, 42, 56, 66, 67, 75, 88}
	// fmt.Println(search.Binary_Search(arr, 32))

	/* Selection Sort */
	// arr := []int{75, 67, 42, 11, 13, 88, 7, 32, 56, 29, 12, 3, 66}
	// fmt.Println(sort.Selection_Sort(arr))

	/* Linked Lists */
	// lL := linear.LinkedList[int]{}

	// lL.AddToBegining(20)

	// lL.AddToEnd(33)
	// lL.AddToEnd(19)
	// lL.AddToEnd(21)
	// lL.AddToEnd(80)
	// lL.AddToEnd(52)
	// lL.AddToEnd(44)

	// _, err = lL.Remove(21)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = lL.Display()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	/* Stacks */
	// stack := linear.Stack[int]{}
	// stack.Push(1)
	// stack.Push(2)
	// stack.Push(3)

	// stack.Pop()
	// _, err = stack.Pop()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = stack.Display()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	/* Queues */
	queue := linear.Queue[int]{}
	queue.Enqueue(4)
	queue.Enqueue(10)
	queue.Enqueue(15)
	queue.Enqueue(23)
	queue.Enqueue(35)

	_, err = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	}

	err = queue.Display()
	if err != nil {
		fmt.Println(err)
	}
}
