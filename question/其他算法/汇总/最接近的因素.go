package 汇总

import "math"

func closestDivisors(num int) []int {

	min := math.MaxInt64
	var res []int
	//
	for j := num + 1; j <= num+2; j++ {
		sqrt := int(math.Sqrt(float64(j)))
		for i := 1; i <= sqrt; i++ {
			otherNum := j / i
			if j%i == 0 {
				if int(math.Abs(float64(otherNum-i))) < min {
					min = int(math.Abs(float64(otherNum - i)))
					res = []int{i, otherNum}
				}
			}
		}
	}
	return res
}
