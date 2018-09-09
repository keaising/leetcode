package integer

import "strings"

type roman struct {
	one string
	five string
}
var romans map[string]roman {
	"sing"
}

func intToRoman(num int) string {
	romans
	ret:=""
	thousand := (int)(num / 1000)
	if thousand > 0 {
		ret = strings.Repeat("M", thousand)
		num = num % 100
	}
	hundrad := (int)(num / 100)
	if hundrad > 0 {
		switch  {
		case hundrad<=3:
			ret += strings.Repeat("C", hundrad)
		case hundrad == 4:
			ret += "CD"
		default: {
			ret += "D"
			ret += strings.Repeat("C", hundrad -5)
		}
		}
	}
}