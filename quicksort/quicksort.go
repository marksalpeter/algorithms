package quicksort

import (
	"math/rand"

	"github.com/marksalpeter/algorithms/sortable"
)

// Sort implements the quick sort algorithm
// O(n^2) Θ(nlogn) Ω(nlogn)
func Sort(s sortable.Sortable) {
	sort(s, 0, s.Len())
}

func sort(s sortable.Sortable, start, end int) {
	// arrays of len 1 are already sorted
	l := end - start
	if end-start < 2 {
		return
	}

	// select a random element as the pivot and move it to the end of the array
	pivot := end - 1
	random := start + (rand.Int() % l)
	s.Swap(random, pivot)

	// move every number less than the pivot before i
	i := start - 1
	for j := start; j < pivot; j++ {
		if s.Less(j, pivot) {
			i++
			s.Swap(i, j)
		}
	}

	// swap the pivot just after i
	i++
	s.Swap(i, pivot)

	// // sort to the left of the pivot
	sort(s, start, i)

	// // sort to the right of the pivot
	sort(s, i+1, end)
}
