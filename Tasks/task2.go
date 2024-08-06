// isPalanderom checks if a given string is a palindrome.
// It returns true if the string is a palindrome, and false otherwise.
func isPalanderom(s string) bool {

	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}




// wordFreqCount counts the frequency of each word in a given string.
// It returns a map where the keys are the words and the values are the frequencies.
func wordFreqCount(s string) map[string]int {

	outMapp := make(map[string]int)
	wordArry := strings.Split(s, " ")

	for _, word := range wordArry {
		outMapp[word]++
	}

	return outMapp

}
