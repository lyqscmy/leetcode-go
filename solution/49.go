package solution

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	xs := make([]string, len(strs))
	for k, v := range strs {
		xs[k] = sortString(v)
	}
	m := make(map[string][]string)
	for k, v := range xs {
		m[v] = append(m[v], strs[k])
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func sortString(s string) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})
	return string(bs)
}
