package main

import "fmt"

func BubbleSort(sli []int) {
	sliceLength := len(sli)

	for i := range sliceLength {
		hasSwapped := false

		for j := 0; j < sliceLength-i-1; j++ {
			if sli[j] > sli[j+1] {
				// sli[j], sli[j+1] = sli[j+1], sli[j]
				temp := sli[j]
				sli[j] = sli[j+1]
				sli[j+1] = temp
				hasSwapped = true
			}
		}

		if !hasSwapped {
			break
		}
	}
}

func main() {
	nums := []int{5, 1, 8, 0, 4}
	fmt.Println("Original array:", nums)
	BubbleSort(nums)
	fmt.Println("Sorted array:", nums)
}
