package sdk

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Go", "Hello, Go!"},
		{"", "Hello, World!"},
	}
	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := Hello(tc.input)
			if got != tc.expected {
				t.Errorf("Hello(%q) = %q; want %q", tc.input, got, tc.expected)
			}
		})
	}
}
