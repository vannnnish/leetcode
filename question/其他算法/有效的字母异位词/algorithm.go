package 有效的字母异位词

func Method(str1, str2 string) bool {
	byteStr1 := []byte(str1)
	byteStr2 := []byte(str2)
	countMap := make(map[byte]int)
	for _, b := range byteStr1 {
		countMap[b] = countMap[b] + 1
	}
	for _, b := range byteStr2 {
		countMap[b] = countMap[b] - 1
	}
	for _, v := range countMap {
		if v != 0 {
			return false
		}
	}
	return true
}
