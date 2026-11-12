package main

import (
	"reflect"
	"testing"
)

func TestCapitalize(t *testing.T) {
	tests := []struct {
		in, want []string
	}{
		{[]string{"zhōng", "zhòng"}, []string{"Zhōng", "Zhòng"}},
		{[]string{"zho1ng"}, []string{"Zho1ng"}},
		{[]string{"ōng"}, []string{"Ōng"}},
		{[]string{"", "a"}, []string{"", "A"}},
	}

	for _, tt := range tests {
		if got := capitalize(tt.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("capitalize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
