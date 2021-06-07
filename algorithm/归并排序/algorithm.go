package 归并排序

import "math"

func mergeSortFromOnline(arr []int) {
	mergeSort(arr)
}
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

// 归并排序
func mergeSortByMyself(arr []int) {
	apart(arr)
}

//
func apart(arr []int) []int {
	length := len(arr)
	if length < 2 {
		//insertSortOptimization(arr)
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge1(mergeSort(left), mergeSort(right))
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

func mergeSortButton2Top(arr []int) {
	var lenth int = len(arr)
	for size := 1; size <= lenth; size += size {
		for i := 0; i+size < lenth; i += 2 * size { //对[i,i+size-1]和[i+size,i+2*size-1]进行归并
			mergeBU(arr, i, i+size-1, int(math.Min(float64(i+2*size-1), float64(lenth-1)))) // arr left mid right 如果i+2*size>n了越界了就取n-1
		}
	}
}

func mergeBU(arr []int, left, mid, right int) {
	// 将要合并的部分做个拷贝
	var tmp []int = make([]int, right-left+1)
	for i, j := left, 0; i <= right; i++ {
		tmp[j] = arr[i]
		j++
	}
	// i做为左半部分的指针 j作为右半部分的指针
	var i, j int = left, mid + 1
	for k := left; k <= right; k++ {
		if i > mid { // 左半部分 已经合入完了将右半部分剩下的 全部合入
			arr[k] = tmp[j-left]
			j++
		} else if j > right { // 右半部分 已经合入完了将左半部分剩下的 全部合入
			arr[k] = tmp[i-left]
			i++
		} else if tmp[i-left] > tmp[j-left] {
			arr[k] = tmp[j-left]
			j++
		} else {
			arr[k] = tmp[i-left]
			i++
		}
	}
}

// 合并俩个有序数组
func merge1(left []int, right []int) []int {
	var result []int
	lenL := len(left)
	lenR := len(right)
	// 从大到小排列
	for ; lenL+lenR > 0; {
		if lenL != 0 && lenR != 0 {
			if left[lenL-1] >= right[lenR-1] {
				result = append(result, left[lenL-1])
				lenL--
			} else {
				result = append(result, right[lenR-1])
				lenR--
			}
		} else if lenL == 0 && lenR != 0 {
			for ; lenR > 0; {
				result = append(result, right[lenR-1])
				lenR--
			}
		} else if lenR == 0 && lenL != 0 {
			for ; lenL > 0; {
				result = append(result, left[lenL-1])
				lenL--
			}
		}
	}
	return result
}
