package 选择排序

func bubbleSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			// 从小到大排序
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func bubbleSortOptimization(nums []int) {
	length := len(nums)
	sortedPosition := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-sortedPosition; j++ {
			// 从小到大排序
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		sortedPosition++
	}
}
