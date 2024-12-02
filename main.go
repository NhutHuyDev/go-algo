package main

import (
	"fmt"

	"github.com/NhutHuyDev/go-algo/ds/linear"
)

func main() {
	var err error

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
	// queue.Enqueue(10)
	// queue.Enqueue(15)
	// queue.Enqueue(23)
	// queue.Enqueue(35)

	value, err := queue.Dequeue()
	// _, _ = queue.Dequeue()
	// _, _ = queue.Dequeue()
	// _, _ = queue.Dequeue()
	// _, _ = queue.Dequeue()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	// err = queue.Display()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	/* Undirected & Unweighted Graph */
	// g := nolinear.NewGraph()

	// Add edges
	// g.AddEdge(1, 2)
	// g.AddEdge(1, 3)
	// g.AddEdge(2, 4)
	// g.AddEdge(2, 5)
	// g.AddEdge(3, 6)
	// g.AddEdge(3, 7)

	// fmt.Println(g.GetAnyNode())

	// g.Display()
}
