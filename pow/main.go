package main

// The task: https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n *= -1
	}

	res, m := 1.0, 1
	for i := 0; i < 32; i++ {
		if n&m > 0 {
			res *= x
		}
		x *= x
		m <<= 1
	}

	return res
}

func main() {
	println(myPow(2.01, 3))
}
