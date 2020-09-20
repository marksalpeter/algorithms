package heap

// NewMinHeap returns a new Heap that re
func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

// MinHeap is a binary min heap implementation
type MinHeap struct {
	list []Element
	size int
}

// PopMin removes the largest element from the heap
func (h *MinHeap) PopMin() interface{} {
	if h.size == 0 {
		return nil
	}

	// get the largest element
	e := h.list[1]

	// copy the last element into the first position
	h.list[1] = h.list[h.size]

	// delete the last elemtn
	h.size--

	// percolate the change down the tree
	h.percolateDown(1)

	// return the previous head
	return e
}

// PeakMin returns the larges element from the heap without removing it
func (h *MinHeap) PeakMin() interface{} {
	return h.list[1]
}

// Insert inserts an element into the heap
func (h *MinHeap) Insert(e Element) {
	if h.size == 0 {
		h.list = []Element{nil}
	}

	// increment the size of the heap
	h.size++

	// add the new item the first open slot
	h.list = append(h.list, e)

	// now percolate the item upwards to its correct position
	h.percolateUp(h.size)
}

func (h *MinHeap) percolateUp(i int) {
	// we're at the root node
	if i == 1 {
		return
	}

	// swap the element with its parent
	pi := i / 2
	if child, parent := h.list[i], h.list[pi]; child.Less(parent) {
		h.list[pi] = child
		h.list[i] = parent
		// percolate up
		h.percolateUp(pi)
	}

}

func (h *MinHeap) percolateDown(i int) {
	// leaf node
	if 2*i > h.size {
		return
	}

	// there is only a left child on the final node
	swapIndex := 2 * i
	if swapIndex == h.size {
		// perform the last swap
		if child, parent := h.list[swapIndex], h.list[i]; child.Less(parent) {
			h.list[swapIndex] = parent
			h.list[i] = child
			return
		}
	}

	// determine wether to go down the left or right node
	if left, right := h.list[swapIndex], h.list[(swapIndex)+1]; right.Less(left) {
		swapIndex++
	}

	// perform a swap if the child node is larger than the parent
	if child, parent := h.list[swapIndex], h.list[i]; child.Less(parent) {
		h.list[swapIndex] = parent
		h.list[i] = child
		//bubble the change down
		h.percolateDown(swapIndex)
	}
}
