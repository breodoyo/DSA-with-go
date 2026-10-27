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
func TestCountOccurrence(t *testing.T) {
	tests := []struct {
		name string
		word string
		target byte
		want int
	}{
		{"repeated character", "banana", 'a', 3},
		{"two occurrences", "hello", 'l', 2},
		{"one occurrence", "hello", 'h', 1},
		{"character absent", "hello", 'z', 0},
		{"empty string", "", 'a', 0},
		{"empty target occurrence", "aaaa", 'a', 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountOccurrence(tt.word, tt.target)

			if got != tt.want {
				t.Errorf("CountOccurrence(%q %q) = %q, want %q", tt.word, tt.target, got, tt.target)
			}
		})
	}
}
func TestContains(t *testing.T) {
	tests := []struct {
		name string
		word string
		target byte
		want bool
	}{
		{"character exists", "hello", 'e', true},
		{"character absent", "hello", 'z', false},
		{"first character", "hello", 'h', true},
		{"last character", "hello", 'o', true},
		{"empty string", "", 'a', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t*testing.T) {
			got := Contains(tt.word, tt.target)
			if got != tt.want {
				t.Errorf("Contains(%q %q) = %t, want %t", tt.word, tt.target, got, tt.want)
			}
		})
	}
}
func TestCountVowels(t *testing.T) {
	tests := []struct {
		word string
		want int
	}{
		{"hello", 2},
		{"banana", 3},
		{"rhythm", 0},
		{"", 0},
		{"aeiou", 5},
	}
	for _, tt := range tests {
		got := CountVowels(tt.word)
		if got != tt.want {
			t.Errorf("countVowels(%q) = %q, want %q", tt.word, got, tt.want)
		}
	}
}
func TestCountConsonants(t *testing.T) {
	tests := []struct {
		word string
		want int
	}{
		{"hello", 3},
   		{"banana", 3},
   		{"rhythm", 6},
    	{"", 0},
    	{"aeiou", 0},
    	{"APPLE", 3},
    	{"HeLLo", 3},
    	{"abc123!", 2},
	}
	for _, tt := range tests {
		got := CountConsonants(tt.word)
		if got != tt.want {
			t.Errorf("CountConsonants(%q) = %d want %d", tt.word, got, tt.want)
		}
	}
}