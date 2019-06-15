package path

import "strings"

func simplifyPath(path string) string {
	s := new(Stack)
	for _, str := range strings.Split(path, "/") {
		switch {
		case !contains(str):
			s.Push(str)
		case str == "..":
			s.Pop()
		case str == "." || str == "":

		}
	}
	return s.toString()
}

type Stack []string

func (s *Stack) Pop() {
	if len(*s) != 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *Stack) Push(str string) {
	(*s) = append((*s), str)
}

func (s *Stack) toString() string {
	if len(*s) == 0 {
		return "/"
	}
	result := ""
	for _, str := range *s {
		result = result + "/" + str
	}
	return result
}

func contains(str string) bool {
	for _, s := range []string{".", "..", ""} {
		if s == str {
			return true
		}
	}
	return false
}
