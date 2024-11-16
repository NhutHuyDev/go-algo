package linear

// Node represents a single node in the linked list.
type Node[T any] struct {
	data T
	next *Node[T]
}
