package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kiransuren/gogrep/search"
	"github.com/kiransuren/gogrep/utils"
)

func main() {
	ignoreArr := []string{`\w*\.git`, `\w*\.exe`}
	ReadDirectories("./", ignoreArr, "echo")
}

// Read file contents and search for pattern using Boyer-Moore (output any matches)
func BoyerMooreSearchFile(pattern string, filename string) {
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

// Format matches by printing the line it was found in and filename
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

// Read directories recursively and find any matches (handle with BoyerMooreSearchFile func)
func ReadDirectories(rootDir string, ignoreArr []string, pattern string) bool {
	files, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Println("")
		return false
	}
	for _, f := range files {

		// Ignore flagged names (whether its a dir or file)
		if utils.TargetContainsRegex(f.Name(), ignoreArr) {
			continue
		}

		if f.IsDir() {
			// Recurse directories
			ReadDirectories(rootDir+f.Name()+"/", ignoreArr, pattern)
		} else {
			// Search for pattern in file using Boyer-Moore
			BoyerMooreSearchFile(pattern, rootDir+f.Name())
		}
	}
	return true
}
