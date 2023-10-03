package binarySearch

func BinarySearch(orderedArray []int, need int) int {
	low := 0
	high := len(orderedArray) - 1

	for low <= high {
		middlePoint := (high + low) / 2
		middleValue := orderedArray[middlePoint]

		if middleValue == need {
			return middlePoint
		} else if need < middleValue {
			high = middlePoint - 1
		} else {
			low = middlePoint + 1
		}
	}

	return -1
}
