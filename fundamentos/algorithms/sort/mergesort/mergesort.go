package main

import "fmt"

/*
Divide y venceras (divide and conquer)
*/

func main() {
	numbers := []int{4, 2, 7}
	sortedNumbers := mergeSort(numbers)
	fmt.Println("Sorted array:", sortedNumbers) // Output: [2 4 7]
}

func mergeSort(arr []int) []int {
	arrayLength := len(arr)

	if arrayLength <= 1 {
		return arr
	}

	middle := arrayLength / 2

	left := arr[:middle]
	right := arr[middle:]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	leftLen := len(left)
	rightLen := len(right)

	newArray := make([]int, 0, leftLen+rightLen)

	i := 0
	j := 0

	for i < leftLen && j < rightLen {
		if left[i] < right[j] {
			newArray = append(newArray, left[i])
			i++
		} else {
			newArray = append(newArray, right[j])
			j++
		}
	}

	for i < leftLen {
		newArray = append(newArray, left[i])
		i++
	}

	for j < rightLen {
		newArray = append(newArray, right[j])
		j++
	}

	// newArray = append(newArray, left[i:]...)
	// newArray = append(newArray, right[j:]...)

	return newArray
}
