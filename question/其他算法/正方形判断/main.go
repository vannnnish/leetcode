package main

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	/*
		取一个点计算 到其他三个点距离
	*/
	p1ToP2 := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	p1ToP3 := (p1[0]-p3[0])*(p1[0]-p3[0]) + (p1[1]-p3[1])*(p1[1]-p3[1])
	p1ToP4 := (p1[0]-p4[0])*(p1[0]-p4[0]) + (p1[1]-p4[1])*(p1[1]-p4[1])

	p2ToP3 := (p2[0]-p3[0])*(p2[0]-p3[0]) + (p2[1]-p3[1])*(p2[1]-p3[1])
	p2ToP4 := (p2[0]-p4[0])*(p2[0]-p4[0]) + (p2[1]-p4[1])*(p2[1]-p4[1])
	p3ToP4 := (p3[0]-p4[0])*(p3[0]-p4[0]) + (p3[1]-p4[1])*(p3[1]-p4[1])
	// 如果p1ToP2恰好是对角线
	if p1ToP3 == p1ToP4 && p1ToP3 != 0 && p1ToP2 != 0 {

		if p2ToP3 == p2ToP4 && p1ToP4 == p2ToP3 && p1ToP2 == p3ToP4 {
			return true
		}
	}
	// 如果p1ToP3是对角线
	if p1ToP2 == p1ToP4 && p1ToP2 != 0 && p1ToP3 != 0 {
		if p2ToP3 == p3ToP4 && p2ToP3 == p1ToP2 && p1ToP3 == p2ToP4 {
			return true
		}

	}
	// 如果p1ToP4是对角线
	if p1ToP2 == p1ToP3 && p1ToP2 != 0 && p1ToP4 != 0 {
		if p2ToP4 == p3ToP4 && p1ToP2 == p2ToP4 && p1ToP4 == p2ToP3 {
			return true
		}
	}
	return false
}
