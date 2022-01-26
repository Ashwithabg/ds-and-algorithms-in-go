package rotation

//Find the Rotation Count in Rotated Sorted array

func findRotationCount(elements []int) int {
	for i := 1; i < len(elements); i++ {
		if elements[i-1] > elements[i] {
			return i
		}
	}
	return 0
}
