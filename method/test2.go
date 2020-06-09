package main

import (
	"fmt"
)

//T t
type T struct {
	int
}

func (t T) test() {
	fmt.Println("类型T方法集包含全部receiver T方法")
}

func (t *T) test2() {
	fmt.Println("类型*T能否包含在T中")
}

func test() {
	t1 := T{1}
	t2 := &t1
	fmt.Printf("t1 is %v\n", t1)
	fmt.Printf("t2 is : %v\n", t2)
	t1.test()
	t2.test()
	t1.test2()
	t2.test2()
}