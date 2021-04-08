package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"github.com/kiransuren/gogrep/search"
)

func main() {
	ignoreArr := []string{`\w*\.git`, `\w*\.exe`}
	ReadDirectories("./", ignoreArr, "echo")
}

func ReadMatchFile(pattern string, filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	matches, err := search.BoyerMooreSearch(pattern, string(data))
	if err != nil {
		log.Fatal(err)
	}
	OutputMatches(data, matches, filename)
}

func OutputMatches(bufferData []byte, matchData []int, bufferName string) {
	for _, pos := range matchData {
		// Loop through ocurrences of pattern in bufferdata
		start, end := pos, pos
		// From pattern ocurrence, walk back until a \r or \n is reached
		for start > 0 && (bufferData[start] != 13 && bufferData[start] != 10) {
			start--
		}
		// From pattern ocurrence, walk forward until a \r or \n is reached
		for end < len(string(bufferData)) && (bufferData[end] != 13 && bufferData[end] != 10) {
			end++
		}
		fmt.Printf(bufferName + ":")
		for x := start + 1; x < end; x++ {
			fmt.Print(string((string(bufferData))[x]))
		}
		fmt.Print("\n")
	}
}

func TargetContainsString(target string, matchingArr []string) bool {
	for _, word := range matchingArr {
		if target == word {
			return true
		}
	}
	return false
}

func TargetContainsRegex(target string, matchingArr []string) bool {
	for _, word := range matchingArr {
		matched, _ := regexp.MatchString(word, target)
		if matched {
			return true
		}
	}
	return false
}

func ReadDirectories(rootDir string, ignoreArr []string, pattern string) bool {
	files, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Println("")
		return false
	}
	for _, f := range files {

		// Ignore flagged names (whether its a dir or file)
		if TargetContainsRegex(f.Name(), ignoreArr) {
			continue
		}

		fmt.Println(rootDir + f.Name())

		if f.IsDir() {
			// Recurse readDirectories
			ReadDirectories(rootDir+f.Name()+"/", ignoreArr, pattern)
		} else {
			// Read file and match with target
			fmt.Println("Reading file: " + rootDir + f.Name())
			ReadMatchFile(pattern, rootDir+f.Name())
		}
	}
	return true
}
