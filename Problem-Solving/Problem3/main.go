package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	record := make([]int, 26)
	for _, c := range s {
		asciiVal := int(c - 'a')
		record[asciiVal]++
	}

	for _, c := range t {
		asciiVal := int(c - 'a')
		if record[asciiVal] == 0 {
			return false
		}
		record[asciiVal]--
	}
	return true
}
