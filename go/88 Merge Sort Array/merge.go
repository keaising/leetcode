package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := 0, 0, 0
	ret := make([]int, len(nums1))
	for m > i && n > j {
		if nums1[i] < nums2[j] {
			ret[k] = nums1[i]
			i++
			k++
		} else {
			ret[k] = nums2[j]
			k++
			j++
		}
	}
	for n > j {
		ret[k] = nums2[j]
		k++
		j++
	}
	for m > i {
		ret[k] = nums1[i]
		k++
		i++
	}
	copy(nums1, ret)
}
