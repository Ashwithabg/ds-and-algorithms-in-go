package main

import "fmt"

func main() {
	str := "hello"
	fmt.Println(isUniqueUsingMap(str))
	fmt.Println(isUniqueUsingArray(str))

	str = "mice"
	fmt.Println(isUniqueUsingMap(str))
	fmt.Println(isUniqueUsingArray(str))
}

func isUniqueUsingArray(str string) bool {
	if len(str) > 128 {
		return false
	}

	values := [128]bool{}
	for _, s := range str {
		if values[s] {
			return false
		}

		values[s] = true
	}

	return true
}

func isUniqueUsingMap(str string) bool {
	values := make(map[rune]bool)

	for _, val := range str {
		_, ok := values[val]
		if ok {
			return false
		}
		values[val] = true
	}

	return true
}
