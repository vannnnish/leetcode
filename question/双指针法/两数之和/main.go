package 两数之和

// 时间复杂度，O(n)  空间复杂度O(n)
func twoSum(numbers []int, target int) []int {
	newResult := make(map[int]int)
	var res []int
	for i, v := range numbers {
		newResult[target-v] = i
	}
	for i := range numbers {
		index, ok := newResult[numbers[i]]
		if ok {
			if i < index {
				res = append(res, i+1, index+1)
				return res
			} else {
				res = append(res, i+1, index+1)
				return res
			}
		}
	}
	return res
}

// 双支阵法
func twoSum1(numbers []int, target int) []int {
	return nil
}

// 返回所查找的元素下标， 如果元素未找到，则返回小于该元素最大值的下标
func binarySearch(numbers []int, target int) int {
	left, right, mid := 0, len(numbers), 0
	for {
		if left > right {
			return mid
		}
		mid = (left + mid) / 2
		if numbers[mid] > target {
			right = mid - 1
		} else if target < numbers[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
}
