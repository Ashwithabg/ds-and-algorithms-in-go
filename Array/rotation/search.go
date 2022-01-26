package rotation

//Search an element in a sorted and rotated array

func getIndexIfExists(numbers []int, element int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if numbers[mid] == element {
		return mid
	}

	if numbers[mid] > element && numbers[start] <= element {
		return getIndexIfExists(numbers, element, start, mid-1)
	}

	if numbers[mid] < element && numbers[end] >= element {
		return getIndexIfExists(numbers, element, mid+1, end)
	}

	if numbers[mid] > element {
		return getIndexIfExists(numbers, element, mid+1, end)
	}

	return getIndexIfExists(numbers, element, start, mid-1)
}
