package main

var memo = make(map[int]int)

func gcd(n int) {

}

func fibonacci(n int) int {
	if n <= 2 {
		return 1
	}

	if val, ok := memo[n]; ok {
		return val
	}

	result := fibonacci(n-1) + fibonacci(n-2)
	memo[n] = result


	return result
}
