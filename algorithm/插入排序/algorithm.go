package 插入排序

func insertSort(nums []int) {
	for i := range nums {
		for j := 0; j < i; j++ {
			// 从小到大排序
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// 找到合适的元素后，就终止循环
func insertSortOptimization(nums []int) {
	for i := range nums {
		for j := i; j > 0; j-- {
			// 从小到大排序
			if nums[j] > nums[j-1] {
				break
			}
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
