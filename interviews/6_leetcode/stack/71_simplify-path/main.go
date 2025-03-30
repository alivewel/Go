package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
    stack := []string{}
    pathArray := strings.Split(path[1:], "/")
	for _, s := range pathArray {
		if s == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if s == "." || s == "" { // s == "" случая не было
			continue
		} else if s != ".." { // s != ".. случая не было
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	// s := "/home/user/Documents/../Pictures"
	// s := "/home/"
	// s := "/home//foo/"
	s := "/../"
	fmt.Println(simplifyPath(s))
}

