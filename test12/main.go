package main

import (
  "fmt"
)

func main()  {
	var fs = [4]func(tt int){}
	a:=make(map[int]int)
	for i:=0;i<4;i++ {
		a[i]=i
		fs[i]= func(tt int) {

			fmt.Println("closure i=",tt,a[tt])
		}
	}
	for x,f:=range fs {
		//fmt.Println(x)
		f(x)
	}
}
