package 归并排序

func MergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	// 分组
	MergeSort(arr, start, mid)
	MergeSort(arr, mid+1, end)
	// 合并
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	tmparr := make([]int, 0)
	var s1, s2 = start, mid + 1
	for s1 <= mid && s2 <= end {
		if arr[s1] > arr[s2] {
			tmparr = append(tmparr, arr[s2])
			s2++
		} else {
			tmparr = append(tmparr, arr[s1])
			s1++
		}
	}

	if s1 <= mid {
		tmparr = append(tmparr, arr[s1:mid+1]...)
	}
	if s2 <= end {
		tmparr = append(tmparr, arr[s2:end+1]...)
	}

	for pos, item := range tmparr {
		arr[start+pos] = item
	}

}
