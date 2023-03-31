package fiboheap

type Node[T any] struct {
	value T
	key float64
	degree int
	mark bool
	parent *Node[T]
	child *Node[T]
	left *Node[T]
	right *Node[T]
}

func NewNode[T any](key float64, value T) *Node[T] {
	return &Node[T]{
		value: value,
		key: key,
		degree: 0,
		mark: false,
		parent: nil,
		child: nil,
		left: nil,
		right: nil,
	}
}