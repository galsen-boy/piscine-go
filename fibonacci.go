package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	}
	if index == 0 {
		return 0
	}
	if index >= 2 {
		return Fibonacci(index-1) + Fibonacci(index-2)
	} else {
		return 1
	}
}
