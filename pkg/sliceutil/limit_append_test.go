package sliceutil

import (
	"reflect"
	"testing"
)

func TestLimitAppend(t *testing.T) {
	tests := []struct {
		limit  int
		slice  [][]byte
		new    []byte
		result [][]byte
	}{
		{
			limit:  4,
			slice:  [][]byte{[]byte("a"), []byte("b"), []byte("c"), []byte("d")},
			new:    []byte("e"),
			result: [][]byte{[]byte("b"), []byte("c"), []byte("d"), []byte("e")},
		},
		{
			limit:  3,
			slice:  [][]byte{[]byte("a"), []byte("b")},
			new:    []byte("c"),
			result: [][]byte{[]byte("a"), []byte("b"), []byte("c")},
		},
	}

	for _, test := range tests {
		got := LimitAppend(test.limit, test.slice, test.new)

		if !reflect.DeepEqual(got, test.result) {
			t.Fatalf("got %q, wanted %q", got, test.result)
		}
	}
}
