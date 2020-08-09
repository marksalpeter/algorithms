package sortable

import (
	"fmt"
	"math/rand"
)

// Sortable is a
type Sortable interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)

	// Ordered returns true if the list is in order
	Ordered() bool

	Log(start, end int)
}

// Ints is a slice of ints
type Ints []int

// Len is the number of elements in the collection.
func (is Ints) Len() int {
	return len(is)
}

// Less reports whether the element with index i should sort before the element with index j.
func (is Ints) Less(i, j int) bool {
	return is[i] < is[j]
}

// Swap swaps the elements with indexes i and j.
func (is Ints) Swap(i, j int) {
	x := is[i]
	is[i] = is[j]
	is[j] = x
}

// Ordered returns true if the list is in order
func (is Ints) Ordered() bool {
	for i := 1; i < is.Len(); i++ {
		if is[i-1] > is[i] {
			return false
		}
	}
	return true
}

// Log returns a subs
func (is Ints) Log(start, end int) {
	fmt.Println(is[start:end])
}

// Random creates a randomized set of ints
func Random(size int) Ints {
	var is []int
	for i := 0; i < size; i++ {
		is = append(is, rand.Int()%100)
	}
	return is
}
