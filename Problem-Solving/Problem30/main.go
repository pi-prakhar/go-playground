package main

/**
Problem :
leetcode -> Climbing Stairs -> easy

Tags :
recursion -> memoization -> dynamic programming -> dp

**/
func solve(n int, dp *[]int) int {
	if n == 1 || n == 0 {
		(*dp)[n] = 1
		return 1
	}
	if (*dp)[n] != 0 {
		return (*dp)[n]
	}
	(*dp)[n] = solve(n-1, dp) + solve(n-2, dp)
	return (*dp)[n]

}
func climbStairs(n int) int {
	dp := make([]int, 46)
	return solve(n, &dp)
}
