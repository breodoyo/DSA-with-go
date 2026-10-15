
package main

import "testing"

func TestLinearSearch(t *testing.T) {
    tests := []struct {
        name    string
        numbers []int
        target  int
        want    int
    }{
        {"first element", []int{5, 8, 11}, 5, 0},
        {"middle element", []int{5, 8, 11}, 8, 1},
        {"last element", []int{5, 8, 11}, 11, 2},
        {"missing target", []int{5, 8, 11}, 20, -1},
        {"empty slice", []int{}, 5, -1},
        {"duplicate values", []int{6, 6, 9}, 6, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := LinearSearch(tt.numbers, tt.target)

            if got != tt.want {
                t.Errorf(
                    "LinearSearch(%v, %d) = %d, want %d",
                    tt.numbers, tt.target, got, tt.want,
                )
            }
        })
    }
}


func TestBinarySearch(t *testing.T) {
    tests := []struct {
        name    string
        numbers []int
        target  int
        want    int
    }{
        {
            name:    "middle target",
            numbers: []int{2, 5, 8, 12, 16, 23, 38, 45},
            target:  12,
            want:    3,
        },
        {
            name:    "first target",
            numbers: []int{2, 5, 8, 12, 16, 23, 38, 45},
            target:  2,
            want:    0,
        },
        {
            name:    "last target",
            numbers: []int{2, 5, 8, 12, 16, 23, 38, 45},
            target:  45,
            want:    7,
        },
        {
            name:    "missing target",
            numbers: []int{2, 5, 8, 12, 16, 23, 38, 45},
            target:  10,
            want:    -1,
        },
        {
            name:    "empty slice",
            numbers: []int{},
            target:  5,
            want:    -1,
        },
        {
            name:    "single element found",
            numbers: []int{7},
            target:  7,
            want:    0,
        },
        {
            name:    "single element missing",
            numbers: []int{7},
            target:  9,
            want:    -1,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := BinarySearch(tt.numbers, tt.target)

            if got != tt.want {
                t.Errorf(
                    "BinarySearch(%v, %d) = %d, want %d",
                    tt.numbers,
                    tt.target,
                    got,
                    tt.want,
                )
            }
        })
    }
}
