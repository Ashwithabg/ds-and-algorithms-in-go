package main

import "fmt"

func main() {
	solve(3, 'A', 'B', 'C')
}

func solve(disk int, source rune, middle rune, destination rune) {
	if disk == 0 {
		fmt.Println("plate ", disk, string(source), "=> ", string(destination))
		return
	}

	solve(disk-1, source, destination, middle)
	fmt.Println("plate ", disk,  string(source), "=> ", string(destination))
	solve(disk-1, middle, source, destination)
}
