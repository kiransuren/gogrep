# GoGrep, a golang grep-like clone

### Building GoGrep
- Dev:
  Run ```sh dev.sh```
  1. Builds project to exe (on Windows) in build folder 
  2. Runs executable with some test arguments
- Prod:
  Run ```sh prod.sh```
  1. Builds project to exe (on Windows) in build folder


### Running GoGrep
- Add Environment Variable to gogrep build folder (or move exe/binary to appropriate folder)
- ```gogrep [OPTIONS] PATTERN [FILE/DIRECTORY]```
- Options
  - ```-r``` : read all files under each directory recursively

### Pipeline/Architecture
1. Start gogrep
    Pass in parameters:
      1. target string/regex
      2. root directory (implicit or explicit)
      3. ignores
2. Recursively read contents of root directory
   1. If item is a directory:
      1. Append name to current root 
      2. Recurse search function with new root directory in a new goroutine
   2. If item is a file:
      1. Read contents line by line and search/match with target string/regex
      2. Output a match
3. Matching data to target
   1. Good old fashion, line by line reading and matching string
   2. Using Boyer-Moore algorithm (same one used by GNU grep)
   3. For regexes, using Go's built-in methods (since they're quite fast)


### Current Support
- GoGrep with Boyer-Moore can only handle ASCII characters (need to deal with "weird" files like .exe, .db, etc.)
- Only looks for string patterns (no regex yet)

### Why GoGrep?
1. I was bored
2. Learned Go last Sunday, wanted to test my knowledge out


