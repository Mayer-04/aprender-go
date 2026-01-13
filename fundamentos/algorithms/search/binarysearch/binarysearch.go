package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 6, 7, 8, 9}
	result := binarySearch(numbers, 9)
	fmt.Println(result)
}

func binarySearch(sli []int, target int) int {
	start := 0
	end := len(sli) - 1

	for start <= end {
		mid := (start + end) / 2

		if sli[mid] == target {
			return mid
		} else if sli[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}
