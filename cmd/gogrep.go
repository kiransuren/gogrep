package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kiransuren/gogrep/search"
)

func main() {
	//ignoreArr := []string{".git"}
	//ReadDirectories("./", ignoreArr)

	data, err := ioutil.ReadFile("search/BoyerMoore.go")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:")
	fmt.Println(data)

	matches := search.BoyerMooreSearch("pattern", string(data))
	OutputMatches(data, matches, "search/BoyerMoore.go")
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

func ArrayContainsString(target string, matchingArr []string) bool {
	for _, word := range matchingArr {
		if target == word {
			return true
		}
	}
	return false
}

func ReadDirectories(rootDir string, ignoreArr []string) bool {
	files, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Println("")
		return false
	}
	for _, f := range files {

		// Ignore flagged names (whether its a dir or file)
		if ArrayContainsString(f.Name(), ignoreArr) {
			continue
		}

		fmt.Println(rootDir + f.Name())

		if f.IsDir() {
			// Recurse readDirectories
			ReadDirectories(rootDir+f.Name()+"/", ignoreArr)
		} else {
			// Read file and match with target
			fmt.Println("Reading file: " + rootDir + f.Name())
		}
	}
	return true
}
