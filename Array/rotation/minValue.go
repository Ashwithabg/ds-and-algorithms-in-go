package rotation

//Find the minimum element in a sorted and rotated array

func findMinValue(elements []int) int {
	return binarySearch(elements, 0, len(elements)-1, elements[0])
}

func binarySearch(elements []int, left, right int, maxSum int) int {
	if left < right {
		mid := (left + right) / 2

		if elements[mid] < maxSum {
			maxSum = elements[mid]
		}

		if elements[left] < elements[mid] && maxSum > elements[left]{
			maxSum = elements[left]
		}

		if elements[left] > elements[mid] {
			return binarySearch(elements, left, mid-1, maxSum)
		}

		if elements[right] < elements[mid] && maxSum > elements[right]{
			maxSum = elements[right]
		}

		if elements[right] < elements[mid] {
			return binarySearch(elements, mid+1, right, maxSum)
		}
	}

	return maxSum
}
