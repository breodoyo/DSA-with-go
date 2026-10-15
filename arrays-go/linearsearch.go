package main

func LinearSearch(number []int, target int) int {
	for i, t := range number {
		if t == target {
			return i
		}
	}
	return -1
}
