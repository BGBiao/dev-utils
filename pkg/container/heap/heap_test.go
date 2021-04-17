package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeapQueue(t *testing.T) {
	h := &IntHeap{2, 1, 5}

	// init a heap with h
	heap.Init(h)
	// push a value to the heap h
	heap.Push(h, 3)
	// heap's feature: root is the minimum
	fmt.Printf("minimum: %d\n", (*h)[0])

	// get the heap value one by one and delete value
	if h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

}
