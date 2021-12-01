package queue_test

import (
	"testing"

	"go-exercises/queue"
)

func TestLIFO_Pop(t *testing.T) {
	t.Run("should return error on empty", func(t *testing.T) {
		lifo := queue.LIFO{}
		_, err := lifo.Pop()
		if err != queue.ErrEmptyStack {
			t.Fail()
		}
	})

	t.Run("should pop last inserted first", func(t *testing.T) {
		lifo := queue.LIFO{}
		lifo.Append("one")
		lifo.Append("two")

		for _, expected := range []string{"two", "one"} {
			v, err := lifo.Pop()
			if err != nil {
				t.Fail()
			}
			if v != expected {
				t.Fail()
			}
		}
	})
}
