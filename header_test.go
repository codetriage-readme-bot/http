package http

import (
	"testing"
	"sort"
	"reflect"

	"github.com/gorilla/http/client"
)

var toHeadersTests = []struct {	
	headers map[string][]string
	expected []client.Header
} {
	{ nil, nil },
	{
		map[string][]string{"Host": []string{"a"}},
		[]client.Header{{"Host", "a"}},
	},
	{
		map[string][]string{"Host": []string{"a", "B"}},
		[]client.Header{{"Host", "B"}, {"Host", "a"}},
	},
	{
		map[string][]string{
			"Host": []string{"a"},
			"Connection": []string{"close"},
		},
		[]client.Header{
			{"Connection", "close"},
			{"Host", "a"},
		},
	},
	{
		map[string][]string{
			"Host": []string{"a", "B"},
			"Connection": []string{"close"},
		},
		[]client.Header{
			{"Connection", "close"},
			{"Host", "B"}, {"Host", "a"},
		},
	},
}

func TestToHeaders(t *testing.T) {
	for _, tt := range toHeadersTests {
		actual := toHeaders(tt.headers)
		sort.Sort(client.Headers(actual))
		if !reflect.DeepEqual(tt.expected, actual) {
			t.Errorf("toHeaders(%v): expected %v, got %v", tt.headers, tt.expected, actual)
		}
	}
}
