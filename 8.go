package main

func countBits(n int) []int {
	bits := make([]int, n+1)

	for i := 0; i <= n; i++ {
		bits[i] = bits[i>>1] + i%2
	}

	return bits
}
