package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		num1   string
		num2   string
		expect string
		no     int
	}{
		// {
		// 	num1:   "4",
		// 	num2:   "6",
		// 	expect: "24",
		// 	no:     1,
		// },
		// {
		// 	num1:   "123",
		// 	num2:   "456",
		// 	expect: "56088",
		// 	no:     1,
		// },
		{
			num1:   "123456789",
			num2:   "123456789",
			expect: "15241578750190521",
			no:     1,
		},
	}
	for _, test := range tests {
		act := multiply(test.num1, test.num2)
		if act != test.expect {
			t.Error(test.no, act, test.expect)
		}
	}

}
