
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
