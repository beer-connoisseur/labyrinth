package application

import (
	"container/heap"
)

type priorityQueue[T any] struct {
	items []*T
	less  func(a, b *T) bool
}

func (pq *priorityQueue[T]) Len() int {
	return len(pq.items)
}

func (pq *priorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}

func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *priorityQueue[T]) Push(x any) {
	pq.items = append(pq.items, x.(*T))
}

func (pq *priorityQueue[T]) Pop() any {
	tmp := pq.items
	elem := tmp[len(tmp)-1]
	pq.items = tmp[:len(tmp)-1]
	return elem
}

type PriorityQueueImpl[T any] struct {
	pq *priorityQueue[T]
}

func NewPriorityQueue[T any](less func(a, b *T) bool) *PriorityQueueImpl[T] {
	q := &priorityQueue[T]{
		less: less,
	}
	heap.Init(q)
	return &PriorityQueueImpl[T]{
		pq: q,
	}
}

func (p *PriorityQueueImpl[T]) PushNode(item *T) {
	heap.Push(p.pq, item)
}

func (p *PriorityQueueImpl[T]) PopNode() *T {
	if p.pq.Len() == 0 {
		return nil
	}
	return heap.Pop(p.pq).(*T)
}

func (p *PriorityQueueImpl[T]) Len() int {
	return p.pq.Len()
}
