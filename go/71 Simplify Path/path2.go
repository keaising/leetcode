package path

import "strings"

func simplifyPath2(path string) string {
	arr := strings.Split(path, "/")
	stack := []string{}

	for _, str := range arr {
		curr := strings.TrimSpace(str)
		if curr == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if curr != "." && len(curr) > 0 {
			stack = append(stack, curr)
		}
	}

	return "/" + strings.Join(stack, "/")
}

