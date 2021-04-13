package main

import "fmt"

func main() {
	input := []int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}
	fmt.Println(8&7 == 0)
	fmt.Println(4 & 3)
	fmt.Println(isPowerOfTwo(8))
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(2))
	fmt.Println(isPowerOfTwo(256))
	fmt.Println(isPowerOfTwo(1024))
	fmt.Println(countPairs(input))
}

func countPairs(deliciousness []int) int {
	res := 0
	for i := range deliciousness {
		for j := range deliciousness {
			if i == j {
				continue
			}
			if isPowerOfTwo(deliciousness[i] + deliciousness[j]) {
				fmt.Println(deliciousness[i], deliciousness[j])
				res++
			}
		}
	}
	return res / 2 % (1000000000 + 7)
}

func isPowOf(v int, pow int) bool {
	for {
		if v == pow {
			return true
		} else if v == 1 {
			return false
		}
		if v%pow != 0 {
			return false
		}
		v = v / pow
	}
}

func isPowerOfTwo(x int) bool {
	return x > 0 && (x&(x-1) == 0)
}
