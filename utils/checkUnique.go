package utils

// check if all the characters in the string are unique
func checkUnique(str string) bool {
	checker := 0
	asciiA := int('a')
	for _, char := range str {
		value := int(char) - asciiA
		if (checker & (1 << value)) > 0 {
			return false
		}
		checker |= (1 << value)
	}
	return true
}
