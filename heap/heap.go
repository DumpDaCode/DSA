package heap

// Max heap are use for sorting and Min heap are used for priority queue

type Heap struct {
	length int
	size   int
	data   []int
}

func Parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return 2*i + 1
}

func Right(i int) int {
	return 2*i + 2
}

func (h *Heap) MaxHeapify(i int) {
	l := Left(i)
	r := Right(i)
	largest := i
	if l < h.size && h.data[l] > h.data[i] {
		largest = l
	}
	if r < h.size && h.data[r] > h.data[largest] {
		largest = r
	}
	if largest != i {
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		h.MaxHeapify(largest)
	}
}

func (h *Heap) MinHeapify(i int) {
	l := Left(i)
	r := Right(i)
	largest := i
	if l < h.size && h.data[l] <= h.data[i] {
		largest = l
	}
	if r < h.size && h.data[r] <= h.data[largest] {
		largest = r
	}
	if largest != i {
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		h.MaxHeapify(largest)
	}
}

func (h *Heap) BuildMaxHeap(arr []int) {
	h.size = len(arr)
	h.length = len(arr)
	h.data = make([]int, len(arr))
	copy(h.data, arr)
	for i := (h.size - 1) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
}

func (h *Heap) BuildMinHeap(arr []int) {
	h.size = len(arr)
	h.length = len(arr)
	h.data = make([]int, len(arr))
	copy(h.data, arr)
	for i := (h.size - 1) / 2; i >= 0; i-- {
		h.MinHeapify(i)
	}
}

func (h *Heap) HeapDelete(i int) {
	if i < h.size {
		if h.data[i] < h.data[h.size-1] {

		}
	}
}
