import (
	"strconv"
	"strings"
)

func getPermutation(n int, k int) string {
	list := []string{}
	for i := 1; i <= n ;i ++ {
		list = append(list, strconv.Itoa(i))
	}
	ans := []string{}
    for len(list) > 0 {
		f := fact(len(list)-1)
		y := k/f
		k = k % f
		ans = append(ans, list[y])
		list = remove(list, y)
	}
	return strings.Join(ans, ",")
}

func remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
