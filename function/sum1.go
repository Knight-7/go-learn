package function

import (
	
)

//Sum1 返回切片的和
func Sum1(args ...int) int  {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

//Factorial 阶乘函数
func Factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * Factorial(n - 1)
}