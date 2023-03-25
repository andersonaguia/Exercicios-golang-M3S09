package highestElement

func HighestElement(arr []int) int {
	highestValue := 0

	for i := 1; i < len(arr); i++ {
		if arr[0] < arr[i] {
			arr[0] = arr[i]
			highestValue = arr[i]
		}
	}

	return highestValue
}
