package function

import (
	
)

//Sum2 第一个参数是切片长度，第二个参数是切片
func Sum2(a int, args ...int) int  {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}