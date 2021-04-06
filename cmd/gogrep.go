package main

import (
	"fmt"
	"os"
)

func main() {
	ignoreArr := []string{".git"}
	ReadDirectories("./", ignoreArr)

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
