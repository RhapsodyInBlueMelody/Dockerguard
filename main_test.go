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

func TestCheckImageTag(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Image without explicit tag",
			input: "nginx",
			want:  TagMissing,
		},
		{
			name:  "Explicit latest tag",
			input: "nginx:latest",
			want:  LatestTag,
		},
		{
			name:  "Specific version tag",
			input: "nginx:1.28",
			want:  VersionTag,
		},
		{
			name:  "Registry port without image tag",
			input: "localhost:5000/myapp",
			want:  TagMissing,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkImageTag(tt.input)

			if got != tt.want {
				t.Errorf("checkImageTag(%q) = %d, want %d",
					tt.input, got, tt.want)
			}
		})
	}
}
