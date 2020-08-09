package mergesort

// Sort implements the merge sort algorithm. Note: we cannot sort in place for merge sort
//  O(nlogn) Θ(nlogn)  Ω(nlogn)
func Sort(ints []int) []int {

	// and array of size 1 is sorted
	l := len(ints)
	if l == 1 {
		return ints
	}

	// calculate the mid point
	m := l / 2

	// sort the left and right halves recursively
	left := Sort(ints[:m])
	right := Sort(ints[m:])

	// merge the two halves
	var merged []int
	var iL, iR int
	lL, lR := len(left), len(right)
	for {
		if iL == lL {
			// at the end of the left array, merge in the rest of the right array
			return append(merged, right[iR:]...)
		} else if iR == lR {
			// at the end of the right array, merge in the rest of the left array
			return append(merged, left[iL:]...)
		} else if left[iL] < right[iR] {
			// if the left element is smaller, add it to merged and increment the left array index
			merged = append(merged, left[iL])
			iL++
		} else {
			// if the right array is smaller =, add it to merged and increment the right array index
			merged = append(merged, right[iR])
			iR++
		}
	}

}
