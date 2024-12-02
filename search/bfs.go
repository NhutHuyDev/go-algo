package search

import (
	"fmt"

	"github.com/NhutHuyDev/go-algo/ds/linear"
	"github.com/NhutHuyDev/go-algo/ds/nonlinear"
)

func BFS(g nonlinear.Graph, start, value int) (bool, []int, error) {
	visited := make(map[int]bool)
	queue := linear.Queue[int]{}
	order := make([]int, 0)

	if !g.Contain(start) {
		return false, order, fmt.Errorf("graph does not contain %v", start)
	}

	queue.Enqueue(start)

	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()

		_, exists := visited[node]
		if exists {
			continue
		}

		if node == value {
			return true, order, nil
		}

		visited[node] = true
		order = append(order, node)

		neighbors, _ := g.GetNeighbors(node)

		for _, neighbor := range neighbors {
			queue.Enqueue(neighbor)
		}
	}

	fmt.Printf("order: %v \n", order)
	return false, order, nil
}
