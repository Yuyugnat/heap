package heap

type BinHeap[T any] struct {
	arr []*Node[T]
}

// return a new Binary Heap containing values of type T
func NewBinHeap[T any]() *BinHeap[T] {
	return &BinHeap[T]{
		arr: []*Node[T]{},
	}
}

// some useful methods
func (heap *BinHeap[T]) swap(a, b int) {
	heap.arr[a], heap.arr[b] = heap.arr[b], heap.arr[a]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

// classic heap methods
func (heap *BinHeap[T]) GetMin() *Node[T] {
	return heap.arr[0]
}

func (heap *BinHeap[T]) Insert(key float64, value T) {
	node := &Node[T]{
		key:   key,
		value: value,
	}

	index := len(heap.arr)
	heap.arr = append(heap.arr, node)

	for index > 0 && heap.arr[index].key < heap.arr[parent(index)].key {
		heap.swap(index, parent(index))
		index = parent(index)
	}
}

func (heap *BinHeap[T]) DecreaseKey(key float64) {
	index := len(heap.arr) - 1
	heap.arr[index].key = key

	for index > 0 && heap.arr[index].key < heap.arr[parent(index)].key {
		heap.swap(index, parent(index))
		index = parent(index)
	}
}

// this method assumes that the heap is already a min heap, and that the node at index i is the only node that is smaller than its children
func (heap *BinHeap[T]) MinHeapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i

	if l < len(heap.arr) && heap.arr[l].key < heap.arr[i].key {
		smallest = l
	}

	if r < len(heap.arr) && heap.arr[r].key < heap.arr[smallest].key {
		smallest = r
	}

	if smallest != i {
		heap.swap(i, smallest)
		heap.MinHeapify(smallest)
	}
}

// this method assumes that the heap is already a min heap, and that the node at index i is the only node that is smaller than its children
func (heap *BinHeap[T]) ExtractMin() *Node[T] {
	if len(heap.arr) == 0 {
		return nil
	}

	if len(heap.arr) == 1 {
		min := heap.arr[0]
		heap.arr = []*Node[T]{}
		return min
	}

	min := heap.arr[0]
	heap.arr[0] = heap.arr[len(heap.arr)-1]
	heap.arr = heap.arr[:len(heap.arr)-1]

	heap.MinHeapify(0)
	return min
}
