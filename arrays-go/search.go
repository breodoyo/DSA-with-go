package main

func LinearSearch(number []int, target int) int {
	for i, t := range number {
		if t == target {
			return i
		}
	}
	return -1
}

func BinarySearch(number []int, target int)  int {
	left := 0
	right := len(number) - 1

	for left <= right {
		mid := left + (right-left)/2
		
		if number[mid] == target {
			return mid
		} else if number[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}