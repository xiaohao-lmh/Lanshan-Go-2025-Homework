package main

import "fmt"

func average(sum int, count int) float64 {
	if count == 0 {
		return 0.0
	}
	return float64(sum) / float64(count)
}
func main() {
	sum := 0
	count := 0

	for {
		fmt.Println("请输入一个整数（输入0时结束）:")
		var num int
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		sum += num
		count++
	}
	fmt.Println("总和是：", sum)
	fmt.Println("个数是：", count)
	grade := average(sum, count)
	fmt.Println("平均成绩为:", grade)
	if grade >= 60 {
		fmt.Println("成绩合格")
	} else {
		fmt.Println("成绩不合格")
	}
}
