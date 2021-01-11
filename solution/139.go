package solution

var memoOfWordBreak = make(map[string]bool)

func wordBreak(s string, wordDict []string) bool {
	for k := range memoOfWordBreak {
		delete(memoOfWordBreak, k)
	}
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}
	return wordBreakImpl(s, wordSet)
}

func wordBreakImpl(s string, wordDict map[string]struct{}) bool {
	if b, ok := memoOfWordBreak[s]; ok {
		return b
	}
	if _, ok := wordDict[s]; ok {
		memoOfWordBreak[s] = true
		return true
	}
	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := wordDict[s[i:]]; ok {
			if wordBreakImpl(s[:i], wordDict) {
				memoOfWordBreak[s] = true
				return true
			}
		}
	}
	memoOfWordBreak[s] = false
	return false
}
