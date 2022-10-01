package stringutil

import (
	"reflect"
	"testing"
)

func TestJoinURL(t *testing.T) {
	tests := []struct {
		name   string
		base   string
		paths  []string
		result string
	}{
		{
			name:   "join url",
			base:   "http://example.com",
			paths:  []string{"foo", "bar"},
			result: "http://example.com/foo/bar",
		},
		{
			name:   "join url with `/`",
			base:   "http://example.com/",
			paths:  []string{"foo/", "/bar"},
			result: "http://example.com/foo/bar",
		},
		{
			name:   "join url with `//`",
			base:   "http://example.com/foo//",
			paths:  []string{"//bar"},
			result: "http://example.com/foo/bar",
		},
		{
			name:   "join url with extension",
			base:   "http://example.com/foo",
			paths:  []string{"bar.go"},
			result: "http://example.com/foo/bar.go",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := JoinURL(test.base, test.paths...)

			if !reflect.DeepEqual(got, test.result) {
				t.Fatalf("got %q, wanted %q", got, test.result)
			}
		})
	}
}
