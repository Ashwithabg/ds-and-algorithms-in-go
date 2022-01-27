package orderStatistics

type minHeap []int

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	(*m) = append((*m), x.(int))
}

func (m *minHeap) Pop() interface{} {
	old := (*m)
	root := old[0]
	updated := old[1:]
	*m = updated
	return root
}
