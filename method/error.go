package main

import (
	"errors"
	"fmt"
	"math"
)

func testmain() {
	var r float64
	fmt.Scanf("%f", &r)
	fmt.Println(test3(r))
	area, err := getCircleArea1(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("area = %f\n", area)
	}
}

func getCircleArea(radius float64) (area float64) {
	if radius < 0 {
		panic("半径不能为负")
	}
	return math.Pi * radius * radius
}

func getCircleArea1(radius float64) (area float64, err error) {
	if radius < 0 {
		err = errors.New("半径不能小于0")
		return
	}
	area = math.Pi * radius * radius
	return
}

func test3(radius float64) float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	return getCircleArea(radius)
}