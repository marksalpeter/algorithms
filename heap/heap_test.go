package heap

import "testing"

type Number int

func (a Number) Less(b interface{}) bool {
	return a < b.(Number)
}

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap()
	for _, n := range []int{10, 5, 15, 7, 5, 40, 1, 33, 4, 7, 9, 0} {
		heap.Insert(Number(n))
	}
	last := heap.PopMax().(Number)
	for {
		t.Log(last)
		if max := heap.PopMax(); max == nil {
			break
		} else if last.Less(max) {
			t.Fatalf(`%d > %d`, max, last)
		} else {
			last = max.(Number)
		}
	}
}
