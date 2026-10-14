package main

import (
	"reflect"
	"testing"
)

func TestFindMax(t *testing.T) {
	numbers := []int{4, 8, 2, 9, 1}

	got := FindMax(numbers)
	want := 9

	if got != want {
		t.Errorf("FindMax() = %d, want %d", got, want)
	}
}

func TestFindMin(t *testing.T) {
	numbers := []int{4, 8, 2, 9, 1}

	got := FindMin(numbers)
	want := 1

	if got != want {
		t.Errorf("FindMin() = %d, want %d", got, want)
	}
}

func TestSum(t *testing.T) {
	numbers := []int{2, 4, 6, 8}

	got := Sum(numbers)
	want := 20

	if got != want {
		t.Errorf("Sum() = %d, want %d", got, want)
	}
}

func TestContains(t *testing.T) {
	numbers := []int{2, 4, 6, 8}

	if !Contains(numbers, 6) {
		t.Error("Contains() = false, want true")
	}

	if Contains(numbers, 5) {
		t.Error("Contains() = true, want false")
	}
}

func TestReverse(t *testing.T) {
	numbers := []int{2, 4, 6, 8}
	want := []int{8, 6, 4, 2}

	Reverse(numbers)

	if !reflect.DeepEqual(numbers, want) {
		t.Errorf("Reverse() = %v, want %v", numbers, want)
	}
}

func TestReverseEmptySlice(t *testing.T) {
	numbers := []int{}

	Reverse(numbers)

	if len(numbers) != 0 {
		t.Errorf("Reverse() on empty slice = %v, want empty slice", numbers)
	}
}