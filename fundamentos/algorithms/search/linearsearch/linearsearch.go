package main

import "fmt"

func main() {
	animals := []string{"lobo", "cangrejo", "elefante", "nutria", "anaconda"}
	target := "nutria"
	result := linearSearch(animals, target)
	fmt.Println("índice del animal:", result)
}

func linearSearch(slice []string, target string) int {
	for i, value := range slice {
		if value == target {
			return i
		}
	}
	return -1
}
