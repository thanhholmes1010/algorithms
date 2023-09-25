package main

import "fmt"

type Heap struct {
	compare_func func(a, b int) bool
	arr          []int
}

func NewHeap(compare_func func(a, b int) bool) *Heap {
	return &Heap{
		compare_func: compare_func,
		arr:          make([]int, 0),
	}
}

func (h *Heap) rebuild() {
	n := len(h.arr)
	p := n / 2
	for p >= 0 {
		l := 2*p + 1
		r := 2*p + 2
		if l < n && h.compare_func(h.arr[l], h.arr[p]) {
			h.arr[l], h.arr[p] = h.arr[p], h.arr[l]
		}

		if r < n && h.compare_func(h.arr[r], h.arr[p]) {
			h.arr[r], h.arr[p] = h.arr[p], h.arr[r]
		}
		p--
	}
}

// Push add all values to the heap, then rebuild for first element alway ensure condition bool
// in compare_func
func (h *Heap) Push(x ...int) {
	for _, v := range x {
		h.arr = append(h.arr, v)
	}
	h.rebuild()
}

// Pop remove the first element from the heap, then rebuild all elements remaining in array
// to ensure alway have next element in first index in compare_func
func (h *Heap) Pop() int {
	last := len(h.arr) - 1
	h.arr[0], h.arr[last] = h.arr[last], h.arr[0]
	v := h.arr[last]
	h.arr = h.arr[:last]
	h.rebuild()
	return v
}

// Peak return the first element in the heap, only look ahead mean get element but not pop, only read
func (h *Heap) Peak() int {
	return h.arr[0]
}

func main() {
	// usage:
	heap := NewHeap(func(a, b int) bool {
		return a < b // mean heap will ensure when you pop alway minium
	})

	heap.Push(1, -10, -6, 30, 2, 0, -1)
	v := heap.Pop() // pop and return minimum element is -10
	fmt.Println("next pop: ", v)
	v = heap.Pop() // pop and return minimum element is -6
	fmt.Println("next pop: ", v)
	v = heap.Pop() // pop and return minimum element is -1
	fmt.Println("next pop: ", v)
}
