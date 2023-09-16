package main

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name   string
		number string
		want   string
	}{
		{
			name:   "Test 1",
			number: "(000)-123-222",
			want:   "000123222",
		},
		{
			name:   "Test 2",
			number: "111-222 333",
			want:   "111222333",
		},
		{
			name:   "Test 3",
			number: "(321) 222 123",
			want:   "321222123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			number := normalize(tt.number)

			if number != tt.want {
				t.Errorf("got %v; want %v", number, tt.want)
			}
		})
	}
}

