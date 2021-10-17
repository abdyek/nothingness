package main

func inArray(needle string, haystack []string) bool {
	for i, size := 0, len(haystack); i < size; i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}
