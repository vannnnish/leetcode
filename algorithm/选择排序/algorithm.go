package 选择排序

func choseSort(nums []int) {
	length := len(nums)
	for i := range nums {
		min := nums[i]
		minPosition := i
		for j := i + 1; j < length; j++ {
			if min > nums[j] {
				min = nums[j]
				minPosition = j
			}
		}
		nums[i], nums[minPosition] = nums[minPosition], nums[i]
	}
}
