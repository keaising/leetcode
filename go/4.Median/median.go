package median

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 0 {
		return float64(findKth(nums1, nums2, n/2)+findKth(nums1, nums2, n/2+1)) / 2.0
	}
	return float64(findKth(nums1, nums2, n/2+1))
}

func findKth(arr1, arr2 []int, k int) int {
	if len(arr1) > len(arr2) {
		return findKth(arr2, arr1, k)
	}
	if len(arr1) == 0 {
		return arr2[k-1]
	}
	if k == 1 {
		return minInt(arr1[0], arr2[0])
	}
	return 0
}

func minInt(a,b int) int {
	if a<b {
		return a
	}
	return b
}
