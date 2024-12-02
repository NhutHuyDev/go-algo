package search

import (
	"reflect"
	"testing"

	"github.com/NhutHuyDev/go-algo/ds/nonlinear"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name          string
		graph         nonlinear.Graph
		start         int
		value         int
		expected      bool
		expectedOrder []int
		shouldFail    bool
	}{
		{
			name: "Value exists in graph with random edge order",
			graph: func() nonlinear.Graph {
				g := nonlinear.NewGraph()
				g.AddEdge(3, 5)
				g.AddEdge(1, 3)
				g.AddEdge(4, 6)
				g.AddEdge(1, 2)
				g.AddEdge(2, 4)
				g.AddEdge(5, 6)
				g.AddEdge(6, 7) // Đỉnh cần tìm
				return *g
			}(),
			start:         1,
			value:         7,
			expected:      true,
			expectedOrder: []int{1, 3, 2, 5, 4, 6}, // Thứ tự duyệt trước khi tìm thấy 7
		},
		{
			name: "Value does not exist in graph with random edge order",
			graph: func() nonlinear.Graph {
				g := nonlinear.NewGraph()
				g.AddEdge(2, 4)
				g.AddEdge(1, 2)
				g.AddEdge(1, 3)
				g.AddEdge(3, 5)
				g.AddEdge(5, 6)
				g.AddEdge(4, 6)
				return *g
			}(),
			start:         1,
			value:         8,
			expected:      false,
			expectedOrder: []int{1, 2, 3, 4, 5, 6}, // Duyệt hết nhưng không tìm thấy
		},
		{
			name: "Disconnected graph with random edge order",
			graph: func() nonlinear.Graph {
				g := nonlinear.NewGraph()
				g.AddEdge(1, 2)
				g.AddEdge(4, 6)
				g.AddEdge(5, 6)
				g.AddEdge(1, 3) // Thành phần 1: {1, 2, 3}
				g.AddEdge(4, 5) // Thành phần 2: {4, 5, 6, 7}
				g.AddEdge(6, 7)
				return *g
			}(),
			start:         1,
			value:         7,
			expected:      false,
			expectedOrder: []int{1, 2, 3}, // Không đến được thành phần chứa 7
		},
		{
			name: "Start node not in graph with random edge order",
			graph: func() nonlinear.Graph {
				g := nonlinear.NewGraph()
				g.AddEdge(3, 4)
				g.AddEdge(4, 5)
				g.AddEdge(1, 2)
				g.AddEdge(5, 6)
				g.AddEdge(6, 7)
				g.AddEdge(2, 3)
				return *g
			}(),
			start:         10, // Start node không tồn tại
			value:         7,
			expected:      false,
			expectedOrder: []int{},
			shouldFail:    true,
		},
		{
			name: "Empty graph",
			graph: func() nonlinear.Graph {
				return *nonlinear.NewGraph()
			}(),
			start:         1,
			value:         2,
			expected:      false,
			expectedOrder: []int{},
			shouldFail:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, order, err := BFS(test.graph, test.start, test.value)
			if test.shouldFail {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != test.expected {
					t.Errorf("Expected result %v, but got %v", test.expected, result)
				}
				if !reflect.DeepEqual(order, test.expectedOrder) {
					t.Errorf("Expected order %v, but got %v", test.expectedOrder, order)
				}
			}
		})
	}
}
