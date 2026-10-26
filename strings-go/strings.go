package main

func FirstCharacter(word string) string {
	if len(word) == 0 {
		return ""
	}
	return string(word[0])
}
func CountCharacter(word string) int {
	return len(word)
}
func LastCharacter(word string) string {
	if len(word) == 0 {
		return ""
	}
	return string(word[len(word)-1])
}
func ReverseString(word string) string {
	if word == "" {
		return ""
	}
	result := ""
	for i := len(word)-1; i >= 0; i-- {
		result += string(word[i])
	}
	return result
}