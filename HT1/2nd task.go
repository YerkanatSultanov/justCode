package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	count := strs[0]

	for i := 0; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], count) {
			count = count[:len(count)-1]
			if len(count) == 0 {
				return ""
			}
		}
	}

	return count
}
