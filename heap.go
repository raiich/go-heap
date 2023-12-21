package heap

type Heap[T any] struct {
	Cmp   func(T, T) int
	items []T
}

func (h *Heap[T]) Push(item T) {
	h.items = append(h.items, item)
	i := len(h.items)
	for {
		parent := i / 2
		if parent > 0 && h.Cmp(h.items[i-1], h.items[parent-1]) < 0 {
			h.items[parent-1], h.items[i-1] = h.items[i-1], h.items[parent-1]
			i = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) Pop() (T, bool) {
	length := len(h.items)
	if length == 0 {
		var zero T
		return zero, false
	}
	ret := h.items[0]
	h.items[0] = h.items[length-1]
	i := 0
	for {
		child, ok := h.child((i+1)*2-1, (i+1)*2, length)
		if ok && h.Cmp(h.items[child], h.items[i]) < 0 {
			h.items[i], h.items[child] = h.items[child], h.items[i]
			i = child
		} else {
			break
		}
	}
	h.shrink(length - 1)
	return ret, true
}

func (h *Heap[T]) child(i, j, length int) (int, bool) {
	if i < length {
		if j < length && h.Cmp(h.items[i], h.items[j]) > 0 {
			return j, true
		} else {
			return i, true
		}
	}
	return 0, false
}

func (h *Heap[T]) shrink(length int) {
	shrink := make([]T, length)
	copy(shrink, h.items)
	h.items = shrink
}
