package main

import (
	"fmt"
)

func selectionSort(input_list []int) []int {
	sorted_list := input_list
	for i := 0; i < len(sorted_list); i++ {
		jMin := i
		for j := i + 1; j < len(sorted_list); j++ {
			if sorted_list[j] < sorted_list[jMin] {
				jMin = j
			}
		}
		if jMin != 1 {
			sorted_list[i], sorted_list[jMin] = sorted_list[jMin], sorted_list[i]
		}
	}
	return sorted_list
}

func main() {
	unsorted_list := []int{8, 2, 56, 8, 74, 1}
	sorted_list := selectionSort(unsorted_list)

	fmt.Println("Sorted List:")
	fmt.Println(sorted_list)
}
