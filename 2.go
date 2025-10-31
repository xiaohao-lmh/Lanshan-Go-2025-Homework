package main

import "fmt"

func CalculationFactory(operation string) func(int, int) int {
	switch operation {
	case "+":
		return func(x, y int) int {
			return x + y
		}
	case "-":
		return func(x, y int) int {
			return x - y
		}
	case "*":
		return func(x, y int) int {
			return x * y
		}
	case "/":
		return func(x, y int) int {
			if y != 0 {
				return x / y
			}
			return 0
		}
	default:
		return nil
	}
} //构建四则运算函数工厂
func main() {
	calculator := CalculationFactory("-") //测试减法函数
	fmt.Println(calculator(2, 2))
	calculator = CalculationFactory("+") //加法函数
	fmt.Println(calculator(2, 2))
	calculator = CalculationFactory("*") //乘法函数
	fmt.Println(calculator(2, 2))
	calculator = CalculationFactory("/") //除法函数
	fmt.Println(calculator(2, 2))
}
