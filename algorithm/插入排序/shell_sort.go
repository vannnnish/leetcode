package 插入排序

func shellSort(arr []int) {
	// 间距
	length := len(arr)
	for gap := length / 2; gap >= 1; {
		// 对分组内元素进行排序
		for i := 0; i < length; i = i + gap {
			for j := i; j > 0; j = j - gap {
				// 从小到大
				if arr[j] > arr[j-1] {
					break
				}
				if arr[j] < arr[j-1] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
		}
		gap = gap / 2
	}
}

func ShellSort(a []int) {
	n := len(a)
	h := 1
	for h < n/3 { //寻找合适的间隔h
		h = 3*h + 1
	}
	for h >= 1 {
		//将数组变为间隔h个元素有序
		for i := h; i < n; i++ {
			//间隔h插入排序
			for j := i; j >= h && a[j] < a[j-h]; j -= h {
				swap(a, j, j-h)
			}
		}
		h /= 3
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}