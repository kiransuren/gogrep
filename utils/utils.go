package utils

import "regexp"

// Checks if target string matches any regex in an array
func TargetContainsRegex(target string, matchingArray []string) bool {
	for _, word := range matchingArray {
		matched, _ := regexp.MatchString(word, target)
		if matched {
			return true
		}
	}
	return false
}

// Checks if target string matches any string in an array
func TargetContainsString(target string, matchingArray []string) bool {
	for _, word := range matchingArray {
		if target == word {
			return true
		}
	}
	return false
}

// Return max of two integers
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
