package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/kiransuren/gogrep/search"
	"github.com/kiransuren/gogrep/utils"
)

type Args struct {
	pattern       string
	rootDirectory string
	isRecursive   bool
}

func main() {

	var wg sync.WaitGroup

	// Array of regexes to ignore while searching
	ignoreArr := []string{`\w*\.git`, `\w*\.exe`}

	isRecursive := flag.Bool("r", false, "Do a recursive search of directories")
	flag.Parse()

	args := Args{
		pattern:       flag.Arg(0),
		rootDirectory: flag.Arg(1),
		isRecursive:   *isRecursive,
	}

	if args.pattern == "" || args.rootDirectory == "" {
		fmt.Println("Missing pattern and/or directory/file argument(s)")
		return
	}

	RunGoGrep(args, ignoreArr, args.rootDirectory, &wg)

	// Wait for go routines to complete
	wg.Wait()
}

// Read file contents and search for pattern using Boyer-Moore (output any matches)
func BoyerMooreSearchFile(pattern string, filename string, wg_ptr *sync.WaitGroup) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	matches, err := search.BoyerMooreSearch(pattern, string(data))
	if err != nil {
		wg_ptr.Done()
		log.Fatal(err)
	}
	OutputMatches(data, matches, filename)
	wg_ptr.Done()
}

// Format matching results in a single file by printing the lines they were found in and filename
func OutputMatches(bufferData []byte, matchData []int, bufferName string) {
	outputString := ""

	// If no matches were found, ignore output routine
	if len(matchData) == 0 {
		return
	}

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
		outputString += bufferName + ":"
		for x := start + 1; x < end; x++ {
			outputString += string((string(bufferData))[x])
		}
		outputString += "\n"
	}
	fmt.Print(outputString + "\n")
}

// Read directories (possibly recursively) and find any matches (handle with BoyerMooreSearchFile func)
func RunGoGrep(args Args, ignoreArr []string, directory string, wg_ptr *sync.WaitGroup) bool {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error ocurred reading directory", err)
		return false
	}
	for _, f := range files {

		// Ignore flagged names (whether its a dir or file)
		if utils.TargetContainsRegex(f.Name(), ignoreArr) {
			continue
		}

		if f.IsDir() {
			// Recurse directories
			if !args.isRecursive {
				continue
			}
			RunGoGrep(args, ignoreArr, directory+f.Name()+"/", wg_ptr)
		} else {
			// Search for pattern in file using Boyer-Moore
			(*wg_ptr).Add(1)
			go BoyerMooreSearchFile(args.pattern, directory+f.Name(), wg_ptr)
		}
	}
	return true
}
