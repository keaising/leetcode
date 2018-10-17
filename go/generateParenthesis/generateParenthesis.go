package generateParenthesis

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []string{"()"}
	}
	if n == 2 {
		return []string{"()()", "(())"}
	}
	m := make(map[string]int)
	for _, pair := range generateParenthesis(n - 1) {
		m["()"+pair]=1
		m[pair+"()"]=1
		m["("+pair+")"]=1
	}
	ret := make([]string, 0)
	for k := range m {
		ret =  append(ret, k)
	}
	fmt.Println("lenght of ret:", len(ret))
	fmt.Println("lenght of map:", len(m))
	for i, r := range ret {
		fmt.Println("iter:", i, r)
	}
	return ret
}
