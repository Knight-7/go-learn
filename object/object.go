package main

import (
	"fmt"
)

//Person 一个人的类
type Person struct {
	name string
	sex string 
	age int
}

//Student 一个学生的类
type Student struct {
	*Person
	id int
	addr string
}

//将person的实例返回一个字符串
func (p *Person) toString() string {
	return fmt.Sprintf("name = %s, sex = %s, age = %v\n", p.name, p.sex, p.age)
}

func (s *Student) toString() string {
	return "hello"
}

func main1 () {
	s1 := Student{&Person{"Knight", "man", 23}, 7, "jx"}
	s2 := &s1
	fmt.Println(s1, s2)
	fmt.Println(s1.name, s2.name)
	fmt.Println(s1.Person.name, s2.Person.name)
	fmt.Println(s1.Person.toString(), s2.toString())
}