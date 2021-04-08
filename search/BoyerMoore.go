package search

import (
	"errors"

	"github.com/kiransuren/gogrep/utils"
)

const NUM_OF_CHARS int = 256

// Preprocess pattern into ASCII character value indexable array (for use in Boyer-Moore Search)
func badCharacterPreprocessing(pattern string, badCharacterArray *[NUM_OF_CHARS]int) error {
	// Defaulting characters to have an occurrence of -1 position (they don't exist)
	for i := 0; i < NUM_OF_CHARS; i++ {
		(*badCharacterArray)[i] = -1
	}

	// Assigning each character the last occurrence position based on pattern string
	for position, currRune := range pattern {
		if int(currRune) > NUM_OF_CHARS {
			return errors.New("ERROR: Pattern string is not ASCII Encodeable")
		}
		(*badCharacterArray)[int(currRune)] = position
	}
	return nil
}

// Boyer-Moore Search Algorithm that uses Bad Character heuristic to speedily search buffer string for pattern
func BoyerMooreSearch(pattern string, buffer string) ([]int, error) {
	var badCharacters [NUM_OF_CHARS]int
	matches := make([]int, 0)
	preprocessErr := badCharacterPreprocessing(pattern, &badCharacters)
	if preprocessErr != nil {
		return matches, preprocessErr
	}
	patternLen := len(pattern)
	searchTextLen := len(buffer)
	stride := 0

	for stride <= searchTextLen-patternLen {
		patternIndex := patternLen - 1
		for patternIndex > -1 && pattern[patternIndex] == buffer[stride+patternIndex] {
			// checking if pattern character matches buffer
			patternIndex--
		}

		if int(buffer[stride+patternIndex]) > NUM_OF_CHARS {
			return matches, errors.New("ERROR: Buffer is not ASCII Encodeable")
		}

		if patternIndex < 0 {
			// MATCH OCCURRED
			matches = append(matches, stride)
			if stride+patternLen < searchTextLen {
				// Get next character in search, based on last occurrence
				// move stride to match
				stride += patternLen - badCharacters[buffer[stride+patternLen]]
			} else {
				// Move one at end of text
				stride += 1
			}
		} else {
			// MISMATCH OCCURRED
			// Essentially calculating the vector the stride should move towards to align
			// with next ocurring character in pattern. If vector is negative (i.e. stride is
			// trying to go backwards), Max will force it to move forward 1 (this can happen
			// if last occurance of letter has already be matched)
			stride += utils.Max(1, patternIndex-badCharacters[buffer[stride+patternIndex]])
		}
	}
	return matches, nil
}
