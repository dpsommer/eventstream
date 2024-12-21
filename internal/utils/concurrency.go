package utils

import (
	"container/heap"
	"sync"
)

type SafeHeap struct {
	sync.Mutex
}

func (h *SafeHeap) Init(i heap.Interface) {
	h.Lock()
	defer h.Unlock()

	heap.Init(i)
}

func (h *SafeHeap) Push(i heap.Interface, x any) {
	h.Lock()
	defer h.Unlock()

	heap.Push(i, x)
}

func (h *SafeHeap) Pop(i heap.Interface) any {
	h.Lock()
	defer h.Unlock()

	return heap.Pop(i)
}
