# GoGrep, a golang grep-like clone

### Pipeline/Architecture
1. Start gogrep
    Pass in parameters:
      a. target string/regex
      b. root directory (implicit or explicit)
      c. ignores
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
   
### Why GoGrep?
1. I was bored
2. Learned Go last Sunday, wanted to test my knowledge out
