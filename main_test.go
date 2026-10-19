package main

import (
	"testing"
)

func TestIsFormInstruction(t *testing.T) {

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "FROM instruction",
			input: "FROM nginx:latest",
			want:  true,
		},
		{
			name:  "USER instruction",
			input: "USER nginx",
			want:  false,
		},
		{
			name:  "Empty Line",
			input: "",
			want:  false,
		},
		{
			name:  "Lowercase FROM",
			input: "from nginx:latest",
			want:  true,
		},
		{
			name:  "Leading space",
			input: "    from nginx:latest",
			want:  true,
		},
		{
			name:  "Comment",
			input: "# FROM nginx:latest",
			want:  false,
		},
		{
			name:  "White space only",
			input: "   ",
			want:  false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := isFromInstruction(tt.input)

			// Check whether got matches the expected result here.
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})

	}
}

func TestUseLatestTag(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "Use Latest Tag",
			input: "FROM nginx:latest",
			want:  true,
		}, {
			name:  "Use Version number tag",
			input: "FROM nginx:2.23.1",
			want:  false,
		}, {
			name:  "Use latest os version",
			input: "FROM nginx:latest-alphine",
			want:  false,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := useLatestTag(tt.input)

			// Check whether got matches the expected result here.
			if got != tt.want {
				t.Errorf("got %t, want %t", got, tt.want)
			}
		})

	}

}
