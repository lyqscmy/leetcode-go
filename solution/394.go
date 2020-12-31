package solution

import (
	"strings"
)

func decodeString(s string) string {
	var sb strings.Builder
	k := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			p := i + 1
			balance := 1
			var q int
			for j := p; j < len(s); j++ {
				if s[j] == '[' {
					balance++
				} else if s[j] == ']' {
					balance--
					if balance == 0 {
						q = j
						break
					}
				}
			}
			t := decodeString(s[p:q])
			for j := 0; j < k; j++ {
				sb.WriteString(t)
			}
			k = 0
			i = q
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			k = k*10 + int(s[i]-'0')
		default:
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}
