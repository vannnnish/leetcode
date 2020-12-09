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



func _quickSort(){

}
