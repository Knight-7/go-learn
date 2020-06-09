package main

import (
	"fmt"
)

type fish struct {
	name string 
}

func showType(a ...interface{}) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("type: %T, value: %v\n", a[i], a[i])
	}
}

func stuinfo(info ...interface{}) {
	stuinfo := make(map[string]interface{})
	stuinfo["name"] = "Knight"
	stuinfo["age"] = 18
	stuinfo["address"] = "JiaXing"
	fmt.Println(stuinfo)
}

func main() {
	showType("1")
	showType(1000000000000000 + 1000000000)
	showType(fish{"泡泡"}, false)
	stuinfo()
	var x interface{}
	x = "knight"
	switch v := x.(type) {
	case int:
		fmt.Printf("x is a int, value = %d\n", v)
	case string:
		fmt.Printf("x is a string, value = %s\n", v)
	case bool:
		fmt.Printf("x is a bool, value = %v", v)
	default:
		fmt.Println("unsupport type!")
	}
}