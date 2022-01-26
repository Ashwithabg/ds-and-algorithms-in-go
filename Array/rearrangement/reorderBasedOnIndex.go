package rearrangement

//Given two integer arrays of same size, “arr[]” and “index[]”, reorder elements in “arr[]”
//according to given index array. It is not allowed to given array arr’s length.

func reorder(arr []int, indices []int) []int {
	temp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		temp[indices[i]] = arr[i]

	}
	return temp
}
