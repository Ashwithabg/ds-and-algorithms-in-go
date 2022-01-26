package rotation

//Given a sorted and rotated array, exitsPairWithSum if there is a pair with a given sum

func exitsPairWithSum(arr []int, sum int) bool {
	high := findHighIndex(arr)
	low := (high + 1) % len(arr)

	for high != low {
		pairSum := arr[low] + arr[high]
		if pairSum == sum {
			return true
		}

		if pairSum < sum {
			low = (low + 1) % len(arr)
		}

		if pairSum > sum {
			high = (high - 1 + len(arr)) % len(arr)
		}
	}

	return false
}

func findHighIndex(numbers []int) int {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return i
		}
	}

	return len(numbers) - 1
}
