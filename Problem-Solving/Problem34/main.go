package main

/**
PROBLEM:
leetcode -> Remove Outermost Parentheses -> easy

TAGS:
string, stack
**/

func removeOuterParentheses(s string) string {
	var st int = 0
	start := 0
	ret := ""
	for end, bracket := range s {
		switch bracket {
		case 40:
			st++
		case 41:
			st--
		default:
			break
		}

		if st == 0 {
			if (end - start) > 1 {
				ret = ret + s[start+1:end]
			}
			start = end + 1
		}
	}
	return ret
}
