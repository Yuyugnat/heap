package heap

type FiboHeap[T any] struct {
	n int
	min *Node[T]
	rootList *Node[T]
}

func NewFiboHeap[T any]() *FiboHeap[T] {
	return &FiboHeap[T]{
		n: 0,
		min: nil,
		rootList: nil,
	}
}

func (heap *FiboHeap[T]) Minimum() *Node[T] {
	return heap.min
}

func (heap *FiboHeap[T]) Insert(key float64, value T) *Node[T] {
	node := NewNode(key, value)
	node.left = node
	node.right = node

	heap.mergeWithRootList(node)

	if heap.min == nil || node.key < heap.min.key {
		heap.min = node
	}

	heap.n++
	return node
}

func (heap *FiboHeap[T]) mergeWithRootList(node *Node[T]) {
	if heap.rootList == nil {
		heap.rootList = node
	} else {
		node.right = heap.rootList
		node.left = heap.rootList.left
		heap.rootList.left.right = node
		heap.rootList.left = node
	}
}

func (heap *FiboHeap[T]) Union(otherHeap *FiboHeap[T]) *FiboHeap[T] {
	n := heap.n + otherHeap.n
	
	min := heap.min
	if otherHeap.min.key < min.key {
		min = otherHeap.min
	}

	newFiboHeap := &FiboHeap[T]{
		n: n,
		min: min,
		rootList: heap.rootList,
	}

	last := otherHeap.rootList.left
	otherHeap.rootList.left = newFiboHeap.rootList.left
	newFiboHeap.rootList.left.right = otherHeap.rootList
	newFiboHeap.rootList.left = last
	newFiboHeap.rootList.left.right = newFiboHeap.rootList

	return newFiboHeap
}