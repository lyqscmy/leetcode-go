package solution

func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	max := 0
	maxCount := 0
	for _, t := range tasks {
		m[t-'A']++
		n := m[t-'A']
		if max == n {
			maxCount++
		} else if max < n {
			max = n
			maxCount = 1
		}
	}
	part := max - 1
	partLength := n - (maxCount - 1)
	emptySlots := part * partLength
	leftTasks := len(tasks) - max*maxCount
	idles := 0
	if emptySlots > leftTasks {
		idles = emptySlots - leftTasks
	}
	return idles + len(tasks)
}
