package main

import "fmt"

func FindMax(numbers []int) int {
	number := numbers[0]

	for _, num := range numbers {
		if num > number {
			number = num
		}
	}
	return number
}

func FindMin(numbers []int) int {
	number := numbers[0]

	for _, num := range numbers {
		if num < number {
			number = num
		}
	}
	return number
}
func Sum(number []int) int {
	total := 0
	for _, ch := range number {
		total += ch
	}
	return total
}

func Contains(numbers []int, target int) bool {
	for _, ch := range numbers {
		if ch == target {
			return true
		}
	}
	return false
}

func Reverse(number []int) {
	right := len(number)-1
	left := 0
	for right > left {
		number[right], number[left] = number[left], number[right]
		
		right--
		left++
	}
	fmt.Println()
}

func main() {
    numbers1 := []int{4, 8, 2, 9, 1}
    numbers2 := []int{7, 3, 9, 2, 6}
	numbers := []int{2, 4, 6, 8}
	numbers3 := []int{4, 7, 2, 9, 5}
	numbers4 := []int{2, 4, 6, 8}
	Reverse(numbers4)
	fmt.Println("Reversed:", numbers4)
    fmt.Println("Maximum:", FindMax(numbers1))
    fmt.Println("Minimum:", FindMin(numbers2))
	fmt.Println("Sum:", Sum(numbers))
	fmt.Println("Contain:", Contains(numbers3, 9))
	
}