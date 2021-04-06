package main

import (
	"fmt"
	"os"
)

func main() {
	readDirectories("./")

}

func readDirectories(rootDir string) bool {
	files, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Println("")
		return false
	}
	for _, f := range files {
		fmt.Println(rootDir + f.Name())
		if f.IsDir() {
			// Recurse readDirectories
			readDirectories(rootDir + f.Name() + "/")
		} else {
			fmt.Println("Reading file: " + rootDir + f.Name())
		}
	}
	return true
}
