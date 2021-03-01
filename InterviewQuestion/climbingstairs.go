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

func iterationClimbingStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	stair := make(map[int]int, n)
	stair[1] = 1
	stair[2] = 2
	res := 0
	for i := 3; i <= n; i++ {
		res = stair[i-1] + stair[i-2]
		stair[i] = res
	}

	return res
}
