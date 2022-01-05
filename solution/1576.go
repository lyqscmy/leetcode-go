package solution

import (
	"bytes"
)

var chars = []byte{'a', 'b', 'c'}

func modifyString(s string) string {
	return string(modifyBytes([]byte(s)))
}

func modifyBytes(s []byte) []byte {
	i := bytes.IndexByte(s, '?')
	if i == -1 {
		return s
	}
	left := getCharOrDash(s, i-1)
	right := getCharOrDash(s, i+1)
	var ch byte
	for _, c := range chars {
		if c == left || c == right {
			continue
		}
		ch = c
	}
	s[i] = ch
	return modifyBytes(s)
}

func getCharOrDash(s []byte, i int) byte {
	if i < 0 || i >= len(s) {
		return '-'
	}
	return s[i]
}
