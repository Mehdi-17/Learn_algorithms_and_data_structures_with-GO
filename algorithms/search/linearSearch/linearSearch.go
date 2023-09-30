package linearSearch

func LinearSearch(input []int, need int) int {
	for i := 0; i < len(input); i++ {
		if input[i] == need {
			return i
		}
	}
	return -1
}
