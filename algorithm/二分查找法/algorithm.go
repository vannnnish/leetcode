package 二分查找法

import (
	"math"
)

func BinarySearch(arr []int, target int) int {
	left, right, mid := 0, len(arr), 0
	for {
		if left > right {
			return -1
		}
		mid = (left + right) / 2
		// 要查找的元素在左边
		if arr[mid] > target {
			right = mid - 1
			continue
		}
		if arr[mid] < target {
			left = mid + 1
			continue
		}
		if arr[mid] == target {
			return mid
		}
	}
}

func BinaryFind(arr []int, k int) int {
	left, right, mid := 1, len(arr), 0
	for {
		// mid向下取整
		mid = int(math.Floor(float64((left + right) / 2)))
		if arr[mid] > k {
			// 如果当前元素大于k，那么把right指针移到mid - 1的位置
			right = mid - 1
		} else if arr[mid] < k {
			// 如果当前元素小于k，那么把left指针移到mid + 1的位置
			left = mid + 1
		} else {
			// 否则就是相等了，退出循环
			break
		}
		// 判断如果left大于right，那么这个元素是不存在的。返回-1并且退出循环
		if left > right {
			mid = -1
			break
		}
	}
	// 输入元素的下标
	return mid
}

func BinarySearch2(arr []int, target int) int {
	left, right, mid := 0, len(arr)-1, (len(arr)-1)/2
	for right >= left {
		if target == arr[mid] {
			return mid
		}
		// 元素在右边
		if target > arr[mid] {
			left = mid+1
			mid = left+(right -left) / 2
			continue
		}
		// 元素在左边
		if target < arr[mid] {
			right = mid-1
			mid = left+(right -left) / 2
			continue
		}
	}
	return -1
}
