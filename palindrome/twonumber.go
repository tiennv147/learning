package main

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	numberDigits := int(math.Log10(float64(x))) + 1
	needDigits := numberDigits / 2
	backwardValue := 0
	for ; needDigits > 0; needDigits-- {
		backwardValue = backwardValue*10 + (x % 10)
		x = x / 10
	}
	if numberDigits%2 != 0 {
		x = x / 10
	}

	return x == backwardValue
}

func main() {
	print(isPalindrome(12221))
}
