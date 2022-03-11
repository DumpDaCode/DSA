package heap

func (h *Heap) HeapSort(arr []int) {
	h.BuildMaxHeap(arr)
	for i := h.length - 1; i >= 1; i-- {
		h.data[i], h.data[0] = h.data[0], h.data[i]
		h.size--
		h.MaxHeapify(0)
	}
}
