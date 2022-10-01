package sliceutil

import (
	"reflect"
	"testing"
)

func TestSliceContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []string
		input string
		want  bool
	}{
		{
			name:  "contains",
			slice: []string{"a", "b", "c"},
			input: "a",
			want:  true,
		},
		{
			name:  "not contains",
			slice: []string{"a", "b", "c"},
			input: "d",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := SliceContains(test.slice, test.input)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("got %v, want %v", got, test.want)
			}
		})
	}
}
