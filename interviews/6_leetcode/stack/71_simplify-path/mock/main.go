package main

import (
	"fmt"
	"strings"
)

// split строки по '/'
// "/home//foo/"
// ["home", "", "foo"]
// result '/' + result

func simplifyPath(path string) string {
    pathArray := strings.Split(path, "/")
    stack := []string{}
    for _, pathCur := range pathArray {
        if pathCur == ".." && len(stack) != 0 {
           stack = stack[:len(stack)-1] 
        }
        if pathCur != "" && pathCur != "." && pathCur != ".." {
            stack = append(stack, pathCur)
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

