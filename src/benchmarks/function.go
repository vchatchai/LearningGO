package benchmarks

func recursive(n int) int {

	if n > 0 {
		return n + recursive(n-1)
	}
	return 0

}

func forloop(n int) int {
	var result int = 0
	for ; 0 < n; n-- {
		result = add(result, n)
	}
	return result
}

func add(x, y int) int {
	return x + y
}
