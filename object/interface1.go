package main

import (
	"fmt"
)

//Mover 一个移动的接口
type Mover interface {
	move()
	run()
}

type dog struct {
	name string
}

func (d dog) move() {
	fmt.Printf("%s can moving.\n", d.name)
}

func (d *dog) run() {
	fmt.Printf("%s is running\n", d.name)
}

func main2() {
	var m Mover = &dog{"旺财"}
	m.move()
	m.run()
}