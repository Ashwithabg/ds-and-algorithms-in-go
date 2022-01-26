package main

import "fmt"

func main() {
	fmt.Printf("%+v \n", rotateUsingTempArray([]int{1, 2, 3, 4, 5}, 3))
	fmt.Printf("%+v \n", rotateOneByOne([]int{1, 2, 3, 4, 5}, 3))
	fmt.Printf("%+v \n", rotateUsingReversal([]int{1, 2, 3, 4, 5}, 3))
	fmt.Printf("%+v \n", blockRotation([]int{1, 2, 3, 4, 5}, 3))
}

func blockRotation(values []int, k int) []int {
	return blockRotationRecur(values, 0, k, len(values))
}

func blockRotationRecur(values []int, startIndex int, k int, len int) []int {
	if k == 0 || k == len {
		return values
	}

	if len-k == k {
		nextBlockIndex := len - k + startIndex
		swap(values, startIndex, nextBlockIndex, k)
	}

	// A > B
	if k < len-k {
		swap(values, startIndex, len-k+startIndex, k)
		blockRotationRecur(values, startIndex, k, len-k)
	}

	//B > A
	if k > len-k{
		swap(values, startIndex, k, len-k)
		blockRotationRecur(values, len-k+startIndex, 2*k-len, k)
	}

	return values
}

func swap(values []int, index1 int, index2 int, k int) {
	for i := 0; i < k; i++ {
		temp := values[index1+i]
		values[index1+i] = values[index2+i]
		values[index2+i] = temp
	}
}

func rotateUsingTempArray(values []int, k int) []int {
	temp := values[:k]
	rotatedArray := values[k:]
	return append(rotatedArray, temp...)
}

func rotateOneByOne(values []int, k int) []int {
	for i := 0; i < k; i++ {
		temp := values[0]

		for j := 0; j < len(values)-1; j++ {
			values[j] = values[j+1]
		}

		values[len(values)-1] = temp
	}

	return values
}

func rotateUsingReversal(values []int, k int) []int {
	end := len(values) - 1

	reverse(values, 0, end)
	reverse(values, 0, end-k)
	reverse(values, k-1, end)
	return values
}

func reverse(values []int, start, end int) {
	for start < end {
		temp := values[start]
		values[start] = values[end]
		values[end] = temp
		start++
		end--
	}
}
