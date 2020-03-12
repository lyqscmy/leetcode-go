package solution

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
	str3 := modOfStrings(str1, str2)
	if str1 == str3 {
		return ""
	}
	if len(str3) == 0 {
		return str2
	}
	return gcdOfStrings(str2, str3)
}

func modOfStrings(str1 string, str2 string) string {
	M := len(str1)
	N := len(str2)
	i := 0
	for ; i+N <= M; i = i + N {
		if str1[i:i+N] != str2 {
			break
		}
	}
	return str1[i:]
}
