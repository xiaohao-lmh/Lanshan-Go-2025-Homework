package main

import "fmt"

func unkownc(numbers []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range numbers {
		counts[num]++
	}
	return counts
}

func main() {
	arr := [4]int{1, 2, 2, 3}
	result := unkownc(arr[:])
	fmt.Println(result)
}
