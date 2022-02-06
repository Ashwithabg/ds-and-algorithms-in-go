package main

import "fmt"

func main() {
	fmt.Println(isAnagram("adobe", "abode"))
	fmt.Println(isAnagram("adobea", "abodee"))
	fmt.Println(isAnagram("ado be", "a bod  e	"))
}

func isAnagram(str1, str2 string) bool {
	str1Map := map[rune]int{}
	for _, letter := range str1 {
		if letter == ' ' {
			continue
		}

		str1Map[letter]++
	}
	fmt.Println(str1Map)

	for _, letter := range str2 {
		if letter == ' ' {
			continue
		}

		val, ok := str1Map[letter]
		if !ok {
			return false
		}

		str1Map[letter] = val - 1
		if str1Map[letter] < 0 {
			return false
		}
	}

	for _, v := range str1Map {
		if v != 0 {
			return false
		}
	}
	fmt.Println(str1Map)

	return true
}
