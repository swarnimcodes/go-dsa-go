package sorting

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
