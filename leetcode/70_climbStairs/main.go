package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	return climbStairs(n - 1) + climbStairs(n - 2)
}