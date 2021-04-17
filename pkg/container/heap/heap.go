package heap

import (
	"container/heap"
)

/*
A heap is a common way to implement a priority queue. To build a priority queue, implement the Heap interface with the (negative) priority as the ordering for the Less method, so Push adds items while Pop removes the highest-priority item from the queue.
*/

type IntHeap []int

// a type must implement the Len(),Less(),Swap() method.
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {

	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// generate a priority queue
type Item struct {
	value    string
	priority int
	index    int // inner index must be updated and maintained by the heap.Interface methods.
}

type PriorityQueue []*Item

// Len(),Less(),Swap()
func (pq PriorityQueue) Len() int {
	return len(pq)
}
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)

}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	// avoid memory leak
	// 及时释放slice 中的内存碎片
	old[n-1] = nil
	// for safety
	item.index = -1
	*pq = old[0 : n-1]
	return item

}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)

}
