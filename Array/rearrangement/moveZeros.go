package rearrangement

//Move all zeroes to end of array

func moveZeros(arr []int) []int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}

	for i := count; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}
