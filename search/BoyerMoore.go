package search

const NUM_OF_CHARS int = 256

func badCharacterPreprocessing(pattern string, badCharacterArray *[NUM_OF_CHARS]int) {
	// Defaulting characters to have an occurrence of -1 position (they don't exist)
	for i := 0; i < NUM_OF_CHARS; i++ {
		(*badCharacterArray)[i] = -1
	}

	// Assigning each character the last occurrence position based on pattern string
	for position, currRune := range pattern {
		(*badCharacterArray)[int(currRune)] = position
	}
}

// Currrently only implements bad character heuristic
func BoyerMooreSearch(pattern string, buffer string) []int {
	var badCharacters [NUM_OF_CHARS]int
	matches := make([]int, 0)
	badCharacterPreprocessing(pattern, &badCharacters)
	patternLen := len(pattern)
	searchTextLen := len(buffer)
	stride := 0

	for stride <= searchTextLen-patternLen {
		patternIndex := patternLen - 1
		for patternIndex > -1 && pattern[patternIndex] == buffer[stride+patternIndex] {
			// checking if pattern character matches buffer
			patternIndex--
		}
		if patternIndex < 0 {
			// Match ocurred
			matches = append(matches, stride)
			stride += patternLen
		} else {
			// There was a mismatch
			if badCharacters[int(buffer[stride+patternIndex])] == -1 {
				// Current search text char does not exist in the pattern, move over the entire word
				// until tail is past current search text char
				stride += patternIndex + 1
			} else {
				// Current search text char exists in the pattern, move word over until this pattern char
				// is right beside the current search text char
				stride += patternLen - 1 - badCharacters[int(buffer[stride+patternIndex])]
			}
		}
	}
	return matches
}
