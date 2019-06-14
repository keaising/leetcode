package main

import "testing"

func TestRotate(t *testing.T) {
	tests := []struct {
		input  *ListNode
		expect *ListNode
		k      int
		no     int
	}{
		{
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			expect: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
			k: 0,
		},
	}

	for _, test := range tests {
		if test.expect != rotateRight(test.input, test.k) {
			t.Error(test.no)
		}
	}

}
