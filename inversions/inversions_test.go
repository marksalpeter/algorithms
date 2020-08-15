package inversions

import (
	"testing"

	"github.com/marksalpeter/algorithms/sortable"
)

func TestInversions(t *testing.T) {
	list := sortable.RandomUnique(10)
	t.Log(list)
	t.Log(Count(list))

	list = []int{1, 3, 5, 2, 4, 6}
	t.Log(list)
	t.Log(Count(list))
}
