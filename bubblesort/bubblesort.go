package bubblesort

import "github.com/marksalpeter/algorithms/sortable"

// Sort implements the bubble sort algorithm
// O(n^2)
func Sort(s sortable.Sortable) {
	swap := -1
	for pass := 0; swap != 0; pass++ {
		swap = 0
		for i, l := 1, s.Len()-pass; i < l; i++ {
			if !s.Less(i-1, i) {
				s.Swap(i-1, i)
				swap++
			}
		}
	}
}
