package application

import (
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a < *b })
	if pq == nil {
		t.Fatal("NewPriorityQueue() returned nil")
	}
	if pq.Len() != 0 {
		t.Errorf("expected empty queue, got length %d", pq.Len())
	}
}

func TestPriorityQueueImpl_Len(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a < *b })
	a, b := 1, 2
	pq.PushNode(&a)
	pq.PushNode(&b)

	if pq.Len() != 2 {
		t.Errorf("Len() = %d, want 2", pq.Len())
	}
}

func TestPriorityQueueImpl_PushPopNode(t *testing.T) {
	pq := NewPriorityQueue(func(a, b *int) bool { return *a < *b })
	a, b, c := 3, 1, 2

	pq.PushNode(&a)
	pq.PushNode(&b)
	pq.PushNode(&c)

	got := pq.PopNode()
	if *got != 1 {
		t.Errorf("PopNode() = %d, want 1", *got)
	}
	got = pq.PopNode()
	if *got != 2 {
		t.Errorf("PopNode() = %d, want 2", *got)
	}
	got = pq.PopNode()
	if *got != 3 {
		t.Errorf("PopNode() = %d, want 3", *got)
	}
}
