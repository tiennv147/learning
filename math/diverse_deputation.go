package math

func factorial(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= i
	}
	return result
}

func combination(k int, n int) int {
	if k == n {
		return 1
	}
	return factorial(n) / ((n - k) * factorial(k))
}

func diverse(m int, w int) int {
	if m == 0 || w == 0 || (m+w) < 3 {
		return 0
	}
	case1 := 0
	case2 := 0
	if w >= 2 {
		case1 = combination(1, m) * combination(2, w)
	}
	if m >= 2 {
		case2 = combination(2, m) * combination(1, w)
	}

	return case1 + case2
}
