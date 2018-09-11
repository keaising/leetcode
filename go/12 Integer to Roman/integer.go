package integer

import "strings"

type roman struct {
	one  string
	five string
}

var relation = []struct {
	num  int
	one  string
	five string
}{
	{
		1000,
		"M",
		"I",
	},
	{
		100,
		"C",
		"D",
	},
	{
		10,
		"X",
		"L",
	},
	{
		1,
		"I",
		"V",
	},
}

func intToRoman(num int) string {
	ret := ""
	for _, r := range relation {
		current := (int)(num / r.num)
		if r.num == 1000 {
			ret += strings.Repeat(r.one, current)
		} else {
			switch {
			case current <= 3:
				ret += strings.Repeat(r.one, current)
			case current == 4:
				ret += r.one + r.five
			default:
				ret += r.five + strings.Repeat(r.one, current-5)
			}
		}
		num = num % r.num
	}
	return ret
}
