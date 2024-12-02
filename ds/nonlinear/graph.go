package nonlinear

import (
	"errors"
	"fmt"
)

type Graph struct {
	adj_matrix map[int][]int
}

func NewGraph() *Graph {
	return &Graph{adj_matrix: make(map[int][]int)}
}

func (g *Graph) AddEdge(v1, v2 int) {
	g.adj_matrix[v1] = append(g.adj_matrix[v1], v2)
	g.adj_matrix[v2] = append(g.adj_matrix[v2], v1)
}

func (g *Graph) GetAnyNode() (int, error) {
	for node := range g.adj_matrix {
		return node, nil
	}
	return 0, errors.New("graph is empty")
}
func (g *Graph) GetNeighbors(node int) ([]int, error) {
	if !g.Contain(node) {
		return []int{}, fmt.Errorf("graph does not contain %v", node)
	}

	return g.adj_matrix[node], nil
}

func (g *Graph) Contain(node int) bool {
	_, exists := g.adj_matrix[node]
	return exists
}

func (g *Graph) Display() {
	for vertex, edges := range g.adj_matrix {
		fmt.Printf("%d -> %v\n", vertex, edges)
	}
}
