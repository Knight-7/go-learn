package main

import "fmt"

type washer interface {
	wash()
}

type dryer interface {
	dry()
}

//WashingMachiner 洗衣机
type WashingMachiner interface {
	washer
	dryer
}

type dryy struct {}

func (d dryy) dry() {
	fmt.Println("甩一甩")
}

type haier struct {
	dryy
}

func (h haier) wash() {
	fmt.Println("洗一洗")
}

func main3() {
	var w1 WashingMachiner = haier{}
	w1.dry()
	w1.wash()
}