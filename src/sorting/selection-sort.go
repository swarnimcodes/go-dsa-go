package main

import (
	"fmt"
)

// sorting the list without creating a new one
func selectionSort(input_list []int) {
	for i := 0; i < len(input_list); i++ {
		jMin := i
		for j := i + 1; j < len(input_list); j++ {
			if input_list[j] < input_list[jMin] {
				jMin = j
			}
		}
		if jMin != i {
			input_list[i], input_list[jMin] = input_list[jMin], input_list[i]
		}
	}
}

func main() {
	unsorted_list := []int{74, 13, 96, 31, 909}
	selectionSort(unsorted_list)

	fmt.Println("Sorted List:")
	fmt.Println(unsorted_list) //Sorted in place
}
