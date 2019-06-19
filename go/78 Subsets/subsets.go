package main

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	for _, v := range nums {
		length := len(result)
		for i := 0; i < length; i++ {
			temp := make([]int, len(result[i]))
			copy(temp, result[i])
			temp = append(temp, v)
			result = append(result, temp)
		}
	}
	return result
}
