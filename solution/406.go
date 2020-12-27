package solution

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	},
	)

	for i, p := range people {
		insert(people, i, p[1], p)
	}
	return people
}

func insert(people [][]int, from, to int, p []int) {
	copy(people[to+1:from+1], people[to:from])
	people[to] = p
}
