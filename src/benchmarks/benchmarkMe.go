package benchmarks

func fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo1(n-1) + fibo1(n-2)
	}
}

func fibo2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibo2(n-1) + fibo2(n-2)
}

func fibo3(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func fibo4(n int) int {
	if n <= 2 {
		return 1
	}
	nm1 := 1
	nm2 := 1
	for i := 3; i <= n; i++ {
		nm1, nm2 = nm2, nm1+nm2
	}

	return nm2
}
