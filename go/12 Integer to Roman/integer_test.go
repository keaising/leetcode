package integer

import "testing"

func TestInteger(t *testing.T){
	tests := []struct {
		num int
		ret string
	} {
		{
			3,
			"III",
		},
		{
			4,
			"IV",
		},
		{
			9,
			"IX",
		},
		{
		58,
		"LVIII",
	},
	{
		1994,
			"MCMXCIV",
	},
	}
	for _, test := range tests {
		act:= intToRoman(test.num)
		if act != test.ret {
			t.Error(test.num, test.ret, act)
		}
	}
}
