package rotation

//Given an array, only rotation operation is allowed on array. We can rotate the array as many times as we want.
//Return the maximum possible summation of i*arr[i].

func findMaxRotation(arr []int) int {
	arrSum := 0
	currVal := 0

	n := len(arr)
	for i := 0; i < n; i++ {
		arrSum = arrSum + arr[i]
		currVal = currVal + (i * arr[i])
	}

	maxSum := currVal
	for j := 1; j < n; j++ {
		currVal = currVal + arrSum - n*arr[n-j]

		if currVal > maxSum {
			maxSum = currVal
		}
	}
	return maxSum
}
