// Package inversions counts inversions in an array while sorting them using a merge sort approach
//
// Usage
//
// This can be useful when implementing collaborative filtering algorithms such as Amazons Reccomendation engine
// The number of inversions quantifies how dissimilar two people purchasing histories are. Sorting for a low number of
// inversions can identify similar users on a platform and a normalized verson of the merged list can be used to display
// reccomendations in order of relevance.
//
// Implementation Details
//
// This implementation is a divide and conquer algorithm very similar to merge sort. As we merge items where the second
// half of the list is larger than the first, we keep track of inversions
//
// O(nlogn) Θ(nlogn)  Ω(nlogn)
//
package inversions

// Count counts the number of inversions in both arrays while sorting and merging boths lists into one
func Count(list []int) (sorted []int, inversions int) {
	// fmt.Println(list)
	l := len(list)
	if l == 1 {
		return list, 0
	}

	// count the inversions in the left and right
	m := l / 2
	left, lc := Count(list[:m])
	right, rc := Count(list[m:])
	inversions = lc + rc

	// merge the two sides, counting the inversions each time the left side is larger than the right side
	iL, iR := 0, 0
	lL, lR := len(left), len(right)

	for {
		if iL == lL {
			return append(sorted, right[iR:]...), inversions
		} else if iR == lR {
			// important: the inversions in the left half of the list were already counted
			// during the last iteration on line 57 where l < r
			return append(sorted, left[iL:]...), inversions
		} else if l, r := left[iL], right[iR]; l < r {
			sorted = append(sorted, l)
			iL++
		} else if l > r {
			sorted = append(sorted, r)
			iR++
			// everything remaining on the left side is greater than this digit on the right side
			// and is therefore, inverted
			inversions += len(left[iL:])
		} else {
			sorted = append(sorted, l, r)
			iR++
			iL++
		}
	}
}
