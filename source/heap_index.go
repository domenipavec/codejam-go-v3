package main

import "container/heap"

type HeapIndex struct {
	data []int
	cmp  func(i, j int) bool
}

func NewHeapIndex(N int, cmp func(i, j int) bool) *HeapIndex {
	h := &HeapIndex{
		data: Range(N),
		cmp:  cmp,
	}
	heap.Init(h)
	return h
}

func (h HeapIndex) Len() int           { return len(h.data) }
func (h HeapIndex) Less(i, j int) bool { return h.cmp(h.data[i], h.data[j]) }
func (h HeapIndex) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *HeapIndex) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.data = append(h.data, x.(int))
}

func (h *HeapIndex) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h HeapIndex) First() int {
	return h.data[0]
}

func (h *HeapIndex) PopFirst() int {
	f := h.First()
	heap.Pop(h)
	return f
}

func (h *HeapIndex) FixFirst() {
	heap.Fix(h, 0)
}
