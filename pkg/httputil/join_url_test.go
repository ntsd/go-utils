package httputil

import (
	"reflect"
	"testing"
)

func TestJoinURL(t *testing.T) {
	tests := []struct {
		base   string
		paths  []string
		result string
	}{
		{
			base:   "http://example.com",
			paths:  []string{"foo", "bar"},
			result: "http://example.com/foo/bar",
		},
		{
			base:   "http://example.com/",
			paths:  []string{"foo/", "/bar"},
			result: "http://example.com/foo/bar",
		},
		{
			base:   "http://example.com/foo//",
			paths:  []string{"//bar"},
			result: "http://example.com/foo/bar",
		},
		{
			base:   "http://example.com/foo",
			paths:  []string{"bar.go"},
			result: "http://example.com/foo/bar.go",
		},
	}

	for _, test := range tests {
		got := JoinURL(test.base, test.paths...)

		if !reflect.DeepEqual(got, test.result) {
			t.Errorf("got %q, wanted %q", got, test.result)
		}
	}
}
