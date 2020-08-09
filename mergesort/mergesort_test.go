package mergesort

import (
	"testing"

	"github.com/marksalpeter/algorithms/sortable"
)

func TestSort(t *testing.T) {
	// this list is not ordered
	list := sortable.Random(10)
	if list.Ordered() {
		t.Fatal(list)
	}

	// sort
	list = sortable.Ints(Sort([]int(list)))

	// the list is ordered
	if !list.Ordered() {
		t.Fatal(list)
	}
}
