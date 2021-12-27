package main

import "fmt"

func main() {
	num := []int{50, 4, 6, 170, 8}
	fmt.Println(sort(num))
}

func sort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := sort(items[:len(items)/2])
	second := sort(items[len(items)/2:])
	return merge(first, second)
}

func merge(left []int, right []int) []int {
	var final []int
	lIndex := 0
	rIndex := 0

	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] < right[rIndex] {
			final = append(final, left[lIndex])
			lIndex++
		} else {
			final = append(final, right[rIndex])
			rIndex++
		}
	}
	for ; lIndex < len(left); lIndex++ {
		final = append(final, left[lIndex])
	}
	for ; rIndex < len(right); rIndex++ {
		final = append(final, right[rIndex])
	}
	return final
}
