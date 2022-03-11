package heap

import (
	"errors"

	"github.com/DumpDaCode/DSA/constants"
)

/*
The heap stores only the keys and not the actual elements. Every time we make changes to keys the
same should be reflected in the actual elements that hold those keys

*/

// This one works for max heap

func (h Heap) HeapMax() int {
	return h.data[0]
}

func (h *Heap) HeapExtractMax() (int, error) {
	if h.size > 1 {
		max := h.data[0]
		h.data[0] = h.data[h.size-1]
		h.size--
		h.MaxHeapify(0)
		return max, nil
	}
	return 1, errors.New("Heap Underflow")
}

func (h *Heap) HeapIncreaseKey(i int, key int) error {
	if h.data[i] > h.data[key] {
		return errors.New("value of key is less than previous value")
	}
	h.data[i] = key
	p := Parent(i)
	for i > 1 && h.data[p] <= h.data[i] {
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
		p = Parent(i)
	}
	return nil
}

func (h *Heap) MaxHeapInsert(key int) {
	h.data = append(h.data, key)
	h.data[h.size] = constants.MinInt
	h.size++
	h.HeapIncreaseKey(h.size-1, key)
}

// This one works for min heap.

func (h Heap) HeapMin() int {
	return h.data[0]
}

func (h *Heap) HeapExtractMin() (int, error) {
	if h.size > 1 {
		max := h.data[0]
		h.data[0] = h.data[h.size-1]
		h.size--
		h.MinHeapify(0)
		return max, nil
	}
	return 0, errors.New("Heap Underflow")
}

func (h *Heap) HeapDecreaseKey(i int, key int) error {
	if h.data[i] < h.data[key] {
		return errors.New("value of key is greater than previous value")
	}
	h.data[i] = key
	p := Parent(i)
	for i > 1 && h.data[p] > h.data[i] {
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
		p = Parent(i)
	}
	return nil
}

func (h *Heap) MinHeapInsert(key int) {
	h.data = append(h.data, key)
	h.data[h.size] = constants.MaxInt
	h.size++
	h.HeapIncreaseKey(h.size-1, key)
}
