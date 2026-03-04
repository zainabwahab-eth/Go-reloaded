package operations

func checkVowel(index int, inputSlice []string) bool {
	vowels := "aeiouAEIOUhH"
	ch := inputSlice[index+1][0]

	for _, w := range vowels {
		if ch == byte(w) {
			return true
		}
	}
	return false

}