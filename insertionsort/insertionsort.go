// package insertionsort
//
// Insertion Sort builds a sorted array one element at a time, shifting
// elements that were previously considered 'sorted' if we need to in
// order to fit those element in the right position.
//
// Psuedocode
//
// 1. The first element of the array is sorted
// 2. For each additional element, j:
//    a. Step through the array
//    b. Reach the first element, i, that is > the element being inserted
//    c. Split the array in half at that point
//    d. Reconstruct the array such that [...i-1, j, i...]
//
// Example
//
//   Sort [3, 2, 5, 1]
//   [3]
//   [2, 3]
//   [2,3,5]
//   [1,2,3,5]
//   * done *
//
package insertionsort

// Sort implements the Insertion Sort Algorithm
func Sort(ints []int) []int {
	var sorted []int
	for _, j := range ints {
		for i, l := 0, len(sorted); i <= l; i++ {
			if isLastElem := i == l; isLastElem {
				// append to the end of the list
				sorted = append(sorted, j)
				break
			} else if sorted[i] > j {
				// add the new element just before the larger element
				tmp := append(sorted[:i], j)
				sorted = append(tmp, sorted[i:]...)
				break
			}
		}
	}
	return sorted
}
