package main

import (
	"fmt"
)

//User A sample user class
type User struct {
	Name string
	Email string
}

func (u User) toString() {
	fmt.Printf("%v: %v\n", u.Name, u.Email)
}

func (u *User) setName(name string) {
	u.Name = name
}

func (u User) setName2(name string) {
	u.Name = name
}

func main(){
	u1 := User{"Knight", "knight2347@163.com"}
	u1.toString()
	u2 := User{"Tina", "1203885207@qq.com"}
	u3 := &u2
	u4 := u2
	u3.toString()
	u1.setName2("Metthew")
	u2.setName("Doris")
	u3.toString()
	u1.toString()
	u4.toString()
}