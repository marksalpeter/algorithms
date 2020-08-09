package selectionsort

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

	// bubble sort
	Sort(list)

	// the list is ordered
	if !list.Ordered() {
		t.Fatal(list)
	}
}
