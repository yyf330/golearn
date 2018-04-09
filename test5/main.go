package  main

import (
	"fmt"
)

//1 接口
type I interface {
	Get() int
	Set(int)
}

//2 定义了一个方法 
type S struct {
	Age int
}

func(s S) Get()int {
	return s.Age
}

func(s *S) Set(age int) {
	s.Age = age
}

//函数传值  值用了接口
func f(i I){
	i.Set(10)
	//fmt.Println(i.Get())
}

func main() {
	s := S{}
	f(&s)  //4
	fmt.Println(s.Get())
}
