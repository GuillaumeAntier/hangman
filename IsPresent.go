package hangman

func IsPresent(letter string, word string, display string) (string, bool) { // Function who verify if the letter is present in the word
	isFind := false
	for i, char := range word {
		if i == len(word)-1 && string(char) == letter {
			display = display[:i] + letter
			isFind = true
			continue
		}
		if string(char) == letter {
			display = display[:i] + letter + display[i+1:]
			isFind = true
		}
	}
	return display, isFind
}

func Contains(inter string, x string) bool { // Function who test if a string is present on another string
	for i := 0; i < len(inter); i++ {
		if x == string(inter[i]) {
			return true
		}
	}
	return false
}
