package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println("1到1000的和:%d\n", sum)
}
