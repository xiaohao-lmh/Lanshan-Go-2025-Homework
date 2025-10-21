package main

import "fmt"

func factorial(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var n int
	fmt.Print("请输入一个正整数：")
	fmt.Scan(&n)

	result := factorial(n)
	fmt.Printf("%d! = %d\n", n, result)
}
