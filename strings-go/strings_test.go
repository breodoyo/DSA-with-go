package main

import (
	"testing"
)

func TestFirstCharacter(t *testing.T) {
    tests := []struct {
        name string
        word string
        want string
    }{
        {
            name: "ordinary word",
            word: "Golang",
            want: "G",
        },
        {
            name: "single character",
            word: "A",
            want: "A",
        },
        {
            name: "empty string",
            word: "",
            want: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := FirstCharacter(tt.word)

            if got != tt.want {
                t.Errorf(
                    "FirstCharacter(%q) = %q, want %q",
                    tt.word,
                    got,
                    tt.want,
                )
            }
        })
    }
}
func TestCountCharacter(t *testing.T) {
	tests := []struct {
		name string
		word string
		want int
	}{
		{"ordinary word", "hello", 5},
		{"two letters", "Go", 2},
		{"single character", "a", 1},
		{"empty string", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountCharacter(tt.word)

			if got != tt.want {
				t.Errorf("CountCharacter(%q) = %d, want %d", tt.word, got, tt.want)
			}
		})
	}
}
func TestLastCharacter(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{"Ordinary word", "Golang", "g"},
		{"Two letters", "Go", "o"},
		{"Single Character", "A", "A"},
		{"Empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastCharacter(tt.word)

			if got != tt.want {
				t.Errorf("LastCharacter(%q) = %q, want %q", tt.word, got, tt.want)
			}
		})
	}
}
func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{"ordinary word", "hello", "olleh"},
		{"two letters", "Go", "oG"},
		{"single character", "A", "A"},
		{"empty string", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseString(tt.word)

			if got != tt.want {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.want, got, tt.word)
			}
		})
	}
}