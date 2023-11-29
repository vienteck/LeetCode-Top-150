package main

import "strings"

func main() {

}

func simplifyPath(path string) string {
	components := strings.Split(path, "/")
	stack := make([]string, 0)

	for _, component := range components {
		if component == "" || component == "." {
			//skip empty components and current directory refernces
			continue
		} else if component == ".." {
			//move up one level by popping from the stack
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, component)
		}
	}
	result := "/" + strings.Join(stack, "/")
	return result
}
