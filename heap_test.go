package heap

import (
	"cmp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHeapInt(t *testing.T) {
	h := &Heap[int]{
		Cmp: cmp.Compare[int],
	}

	h.Push(100)
	h.Push(0)
	h.Push(9999)
	h.Push(0)
	h.Push(9999)
	h.Push(100)

	assert.Equal(t, 0, pop(t, h))
	assert.Equal(t, 0, pop(t, h))
	assert.Equal(t, 100, pop(t, h))
	assert.Equal(t, 100, pop(t, h))
	assert.Equal(t, 9999, pop(t, h))
	assert.Equal(t, 9999, pop(t, h))
}

func TestHeapBook(t *testing.T) {
	h := &Heap[book]{
		Cmp: func(a, b book) int {
			if a.at.Before(b.at) {
				return -1
			} else if a.at.After(b.at) {
				return 1
			} else {
				return 0
			}
		},
	}

	h.Push(book{at: time.Unix(100, 100)})
	h.Push(book{at: time.Unix(100, 0)})
	h.Push(book{at: time.Unix(100, 10)})
	h.Push(book{at: time.Unix(0, 10)})
	h.Push(book{at: time.Unix(0, 0)})
	h.Push(book{at: time.Unix(0, 20)})
	h.Push(book{at: time.Unix(9999, 999)})
	h.Push(book{at: time.Unix(9999, 0)})
	h.Push(book{at: time.Unix(9999, 99)})
	h.Push(book{at: time.Unix(0, 0)})
	h.Push(book{at: time.Unix(9999, 999)})
	h.Push(book{at: time.Unix(100, 10)})

	assert.Equal(t, book{at: time.Unix(0, 0)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(0, 0)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(0, 10)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(0, 20)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(100, 0)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(100, 10)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(100, 10)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(100, 100)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(9999, 0)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(9999, 99)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(9999, 999)}, pop(t, h))
	assert.Equal(t, book{at: time.Unix(9999, 999)}, pop(t, h))
}

func pop[T any](t *testing.T, h *Heap[T]) T {
	t.Helper()
	v, ok := h.Pop()
	if !ok {
		t.Fatal("pop failed")
	}
	return v
}

type book struct {
	at time.Time
}
