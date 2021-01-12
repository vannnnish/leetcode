package 非递减数列

import "fmt"

func checkPossibility(nums []int) bool {
	length := len(nums)
	var decreasingCount int
	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			if i == 0 || nums[i-1] <= nums[i+1] {
				decreasingCount++
				nums[i] = nums[i+1]
			} else {
				decreasingCount++
				nums[i+1] = nums[i]
			}
		}
		fmt.Println(decreasingCount)
		if decreasingCount > 1 {
			return false
		}
	}
	fmt.Println(nums)
	for i := 0; i < length-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
