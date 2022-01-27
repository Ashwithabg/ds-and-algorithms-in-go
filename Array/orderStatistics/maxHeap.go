package orderStatistics

type maxHeap []int

func (n maxHeap) Len() int {
	return len(n)
}

func (n maxHeap) Less(i, j int) bool {
	return n[i] > n[j]
}

func (n maxHeap) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n *maxHeap) Push(x interface{}) {
	*n = append(*n, x.(int))
}

func (n *maxHeap) Pop() interface{} {
	old := *n
	x := old[0]
	updated := old[1:]
	*n = updated
	return x
}
