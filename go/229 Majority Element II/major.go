package main

func majorityElement(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}
	n1, n2, c1, c2 := 0, 0, 0, 0
	for _, n := range nums {
		if n == n1 {
			c1++
		} else if n == n2 {
			c2++
		} else if c1 == 0 {
			n1 = n
			c1++
		} else if c2 == 0 {
			n2 = n
			c2++
		} else {
			c2--
			c1--
		}
	}
	c1, c2 = 0, 0
	for _, n := range nums {
		if n == n1 {
			c1++
		} else if n == n2 {
			c2++
		}
	}
	ans := []int{}
	if c1 > len(nums)/3 {
		ans = append(ans, n1)
	}
	if c2 > len(nums)/3 {
		ans = append(ans, n2)
	}
	return ans
}
