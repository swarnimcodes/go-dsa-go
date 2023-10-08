package sorting

func bubbleSort(arr []int) {
	swapped := false
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			if swapped == false {
				break
			}
		}
	}
}
