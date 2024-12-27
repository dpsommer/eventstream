package utils

// Adapted from the container/heap priority queue example:
// https://pkg.go.dev/container/heap#example-package-PriorityQueue
//
// we don't need to implement update, so the index field is unnecessary.
// refactored using generics to avoid tight coupling with graph impl

type Prioritizable interface {
	Priority() int64
}

// A PriorityQueue implements heap.Interface and holds Prioritizable types
type PriorityQueue[T Prioritizable] []*T

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// Implement as a min-heap, so use less than here
	return (*pq[i]).Priority() < (*pq[j]).Priority()
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	el := x.(*T)
	*pq = append(*pq, el)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	el := old[n-1]
	old[n-1] = nil // don't stop the GC from reclaiming the element eventually
	*pq = old[0 : n-1]
	return el
}
