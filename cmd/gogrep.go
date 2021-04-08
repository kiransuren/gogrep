package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kiransuren/gogrep/search"
	"github.com/kiransuren/gogrep/utils"
)

type Args struct {
	pattern       string
	rootDirectory string
	isRecursive   bool
}

func main() {

	isRecursive := flag.Bool("r", false, "Do a recursive search of directories")
	flag.Parse() // This will parse all the arguments from the terminal

	args := Args{
		pattern:       flag.Arg(0),
		rootDirectory: flag.Arg(1),
		isRecursive:   *isRecursive,
	}

	if args.pattern == "" || args.rootDirectory == "" {
		fmt.Println("Missing pattern and/or directory/file argument(s)")
		return
	}

	ignoreArr := []string{`\w*\.git`, `\w*\.exe`}
	RunGoGrep(args, ignoreArr, args.rootDirectory)
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

// Read directories (possibly recursively) and find any matches (handle with BoyerMooreSearchFile func)
func RunGoGrep(args Args, ignoreArr []string, directory string) bool {
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
			RunGoGrep(args, ignoreArr, directory+f.Name()+"/")
		} else {
			// Search for pattern in file using Boyer-Moore
			BoyerMooreSearchFile(args.pattern, directory+f.Name())
		}
	}
	return true
}
