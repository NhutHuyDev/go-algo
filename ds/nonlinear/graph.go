package nolinear

import "fmt"

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

func (g *Graph) Display() {
	for vertex, edges := range g.adj_matrix {
		fmt.Printf("%d -> %v\n", vertex, edges)
	}
}
