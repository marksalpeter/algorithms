package selectionsort

import "github.com/marksalpeter/algorithms/sortable"

// Sort implements the bubble sort algorithm
// O(n^2) Ω(n^2)
func Sort(s sortable.Sortable) {
	for i, l := 0, s.Len(); i < l; i++ {
		for j := i; j < l; j++ {
			if !s.Less(i, j) {
				s.Swap(i, j)
			}
		}
	}
}
