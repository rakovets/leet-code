package main

func longestPalindrome(s string) string {
	result := ""
	if (len(s) != 0) {
		result = s[0:1]
	}
	

	leftBorder := 0
	for leftBorder <= len(s) - 1 {
		rigthBorder := leftBorder + 1
		for rigthBorder <= len(s) {
			if (isPalindrome(s[leftBorder:rigthBorder]) && len(s[leftBorder:rigthBorder]) > len(result)) {
				result = s[leftBorder:rigthBorder]
			}
			rigthBorder++
		}
		leftBorder++
	} 
	return result
}

func isPalindrome(s string) bool {
	i := 0
	for i < len(s) / 2 {
		if (s[i] != s[len(s) -1 - i]) {
			return false
		}
		i++
	}
	return true
}
