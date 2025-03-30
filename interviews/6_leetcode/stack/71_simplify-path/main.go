package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
    stack := []string{}
    pathArray := strings.Split(path[1:], "/")
	// проверить что будет если в пути // или ///
	for _, s := range pathArray {
		if s == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else if s == "." {
			continue
		} else {
			stack = append(stack, s)
		}
	}
	// fmt.Println(pathArray)
	result := "/" + strings.Join(stack, "/")
	return result
}

func main() {
	s := "/home/user/Documents/../Pictures"
	fmt.Println(simplifyPath(s))
}

