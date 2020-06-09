package test

//Add 相加函数
func Add(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}