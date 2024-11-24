package main

func firstUniqChar(s string) int {
	length := len(s)
	tracker := make(map[byte]struct{})

	for index1 := 0; index1 < length; index1++ {
		if _, ok := tracker[s[index1]]; ok {
			continue
		} else {
			tracker[s[index1]] = struct{}{}
		}
		for index2 := index1; index2 < length; index2++ {
			if index2 == length-1 {
				return index1
			}

			if s[index2+1] == s[index1] {
				break
			}

		}

	}

	return -1
}
