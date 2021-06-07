package 快速排序

import "fmt"

func Method(array []int, begin, end int, mark string) {
	var i, j int
	if begin < end {
		i = begin + 1 // 将array[begin]作为基准数，因此从array[begin+1]开始与基准数比较！
		j = end       // array[end]是数组的最后一位

		for {
			if i >= j {
				break
			}
			if array[i] > array[begin] {
				array[i], array[j] = array[j], array[i]
				j = j - 1
			} else {
				i = i + 1
			}
		}

		/* 跳出while循环后，i = j。
		 * 此时数组被分割成两个部分  -->  array[begin+1] ~ array[i-1] < array[begin]
		 *                           -->  array[i+1] ~ array[end] > array[begin]
		 * 这个时候将数组array分成两个部分，再将array[i]与array[begin]进行比较，决定array[i]的位置。
		 * 最后将array[i]与array[begin]交换，进行两个分割部分的排序！以此类推，直到最后i = j不满足条件就退出！
		 */
		if array[i] >= array[begin] { // 这里必须要取等“>=”，否则数组元素由相同的值时，会出现错误！
			i = i - 1
		}

		array[begin], array[i] = array[i], array[begin]
		fmt.Printf("%s>%v,%d,%v\n", mark, array[begin:i], array[i], array[j:end])
		Method(array, begin, i, mark+"--")
		Method(array, j, end, mark+"--")
	}
}

func quickSort(arr []int, start, end int) {
	var i, j int
	if start < end {
		i = start + 1
		j = end
		for {
			if i >= j {
				break
			}
			if arr[i] > arr[start] {
				arr[i], arr[j] = arr[j], arr[i]
				j--
			} else {
				i++
			}
		}
		if arr[i] >= arr[start] {
			i--
		}
		arr[i], arr[start] = arr[start], arr[i]
		quickSort(arr, start, i)
		quickSort(arr, j, end)
	}

}

func quickSort1(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func quickSort3(arr []int) {
	sort(arr, 0, len(arr)-1)
}

// start 表示起始小标， end 表示结尾下标
// 以起始点为标记点进行排序
func sort(arr []int, start, end int) {
	var left, right int
	if start < end {
		left, right = start+1, end
		for {
			if left >= right {
				break
			}
			// 从小到大排序
			if arr[left] > arr[start] {
				arr[left], arr[right] = arr[right], arr[left]
				right = right - 1
			} else {
				left = left + 1
			}
		}
		if arr[left] >= arr[start] { // 这里必须要取等“>=”，否则数组元素由相同的值时，会出现错误！
			left = left - 1
		}

		arr[start], arr[left] = arr[left], arr[start]

		sort(arr, start, left)
		sort(arr, right, end)
	}
}

func quickSort4(arr []int) {
	sortByMiddle(arr, 0, len(arr)-1)
}

// 以中间点为标记进行
// start 表示起始小标， end 表示结尾下标
// 以起始点为标记点进行排序
func sortByMiddle(arr []int, start, end int) {
	var left, right int
	if start < end {
		left, right = start+1, end
		key := arr[(start+end)/2]
		arr[start], arr[(start+end)/2] = arr[(start+end)/2], arr[start]
		for {
			if left >= right {
				break
			}
			// 从小到大排序
			if arr[left] > key {
				arr[left], arr[right] = arr[right], arr[left]
				right--
			} else {
				left++
			}
		}
		if arr[left] >= key {
			left--
		}
		arr[left], arr[start] = arr[start], arr[left]
		sortByMiddle(arr, start, left)
		sortByMiddle(arr, right, end)
	}
}

func quickSort5(arr []int) {
	threeWaysSort(arr, 0, len(arr)-1)
}

// 以中间点为标记进行
// start 表示起始小标， end 表示结尾下标
// 三路快排
func threeWaysSort(arr []int, start, end int) {
	// i: 当前处理的元素
	// lv: 小于被处理元素的下标
	// gv: 大于标处理元素的下标
	var i, lv, gv int
	if start < end {
		i, lv, gv = start+1, start+1, end
		key := arr[(start+end)/2]
		arr[start], arr[(start+end)/2] = arr[(start+end)/2], arr[start]
		for {
			if i >= gv {
				if arr[i] < key {
					arr[i], arr[lv] = arr[lv], arr[i]
				}
				break
			}
			// 从小到大排序
			// 如果该元素比被比较的元素小，那么放到左边
			if arr[i] > key {
				arr[i], arr[gv] = arr[gv], arr[i]
				gv--
			} else if arr[i] < key {
				arr[i], arr[lv] = arr[lv], arr[i]
				i++
				lv++
			} else {
				i++
			}
		}
		if arr[lv] == key {
			lv--
		}
		arr[lv], arr[start] = arr[start], arr[lv]
		threeWaysSort(arr, start, lv-1)
		threeWaysSort(arr, gv, end)
	}
}
