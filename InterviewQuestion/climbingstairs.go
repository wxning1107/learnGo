package main

func recursiveClimbingStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return recursiveClimbingStairs(n-1) + recursiveClimbingStairs(n-2)
}

func iterationClimbingStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	ret := 0
	pre := 2
	prepre := 1
	for i := 3; i <= n; i++ {
		ret = pre + prepre
		prepre = pre
		pre = ret
	}

	return ret
}
