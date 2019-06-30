package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := 0
	if len(nums) < 3 {
		for _, n:= range nums {
			ans += n
		}
		return ans
	}

	ans = nums[0]+nums[1]+nums[2]
	for i:=0;i<=len(nums)-3;i++ {
		j:=i+1
		k :=len(nums)-1
		for j<k{
			sum:= nums[i]+nums[j]+nums[k]
			if math.Abs(float64(sum - target)) < math.Abs(float64(ans -target)) {
				ans = sum
				if ans == target {
					return ans
				}
			}
			if sum > target {
				k--
			} else {
					j++
			}
		}
	}
	return ans
}