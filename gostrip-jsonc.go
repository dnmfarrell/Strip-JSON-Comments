package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inStr, inMLComment bool
	for scanner.Scan() {
		var inSLComment, escape, star, slash bool
		line := scanner.Bytes()
		for i, c := range line {
			if inStr {
				if !escape {
					if c == '"' {
						inStr = false
					} else if c == '\\' {
						escape = true
					}
				} else {
					escape = false
				}
			} else if inMLComment {
				if star {
					if c == '/' {
						inMLComment = false
					}
					star = false
				} else if c == '*' {
					star = true
				}
				line[i] = ' '
			} else if inSLComment {
				line[i] = ' '
			} else if slash {
				if c == '/' {
					inSLComment = true
					line[i] = ' '
					line[i-1] = ' '
				} else if c == '*' {
					inMLComment = true
					line[i] = ' '
					line[i-1] = ' '
				}
				slash = false
			} else if c == '/' {
				slash = true
			} else if c == '"' {
				inStr = true
			}
		}
		fmt.Println(string(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
