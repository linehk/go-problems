package printListInReversedOrder

import (
	"container/list"
	"testing"
)

func TestPrintListInReversedOrder(t *testing.T) {
	l := list.New()
	input := "12345"
	for _, v := range input {
		l.PushBack(v)
		t.Log()
	}
}
