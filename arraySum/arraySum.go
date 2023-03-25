package arraySum

func ArraySum(arr []int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}
