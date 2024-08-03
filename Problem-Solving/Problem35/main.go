package main

func reverseWords(s string) string {
	start := 0
	length := len(s)
	ret := ""
	for end := 0; end < length; end++ {
		if string(s[end]) == " " && start != end {
			ret = s[start:end] + " " + ret
			start = end + 1
		} else if string(s[end]) == " " && start == end {
			start++
		}
	}
	if start < length {
		ret = s[start:length] + " " + ret
	}
	return ret[:len(ret)-1]
}
