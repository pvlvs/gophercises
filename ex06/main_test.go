package main

import "testing"

func TestCamelcase(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want int32
	}{
		{
			name: "Sample Input 1",
			val:  "saveChangesInTheEditor",
			want: 5,
		},
		{
			name: "Sample Input 2",
			val:  "helloWorld",
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := camelcase(tt.val)

			if val != tt.want {
				t.Errorf("got %v; want %v", val, tt.want)
			}
		})
	}
}

type input struct {
	length int
	s      string
	k      int32
}

func TestCaeserCipher(t *testing.T) {

	tests := []struct {
		name  string
		input input
		want  string
	}{
		{
			name: "Sample Input 1",
			input: input{
				length: 11,
				s:      "middle-Outz",
				k:      2,
			},
			want: "okffng-Qwvb",
		},
		{
			name: "Sample Input 2",
			input: input{
				length: 11,
				s:      "middle-Yutz",
				k:      4,
			},
			want: "qmhhpi-Cyxf",
		},
		{
			name: "Sample Input 3",
			input: input{
				length: 3,
				s:      "XyZ",
				k:      2,
			},
			want: "ZaB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := caesarCipher(tt.input.s, tt.input.k)

			if val != tt.want {
				t.Errorf("got %v; want %v", val, tt.want)
			}
		})
	}
}
