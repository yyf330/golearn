package main

import (
	"fmt"
)


type testBuilder struct {
	a int
	b int
	c string
	d string
}

func (b *testBuilder) to(tt int) *testBuilder {
	b.a = tt
	return b
}


func (b *testBuilder) find(tt int) *testBuilder {
	b.b = tt
	return b
}

func (b *testBuilder) come(tt string) *testBuilder {
	b.c = tt
	return b
}

func (b *testBuilder) admin(tt string) *testBuilder {
	b.d = tt
	return b
}

func main(){
	ws := new(testBuilder)
	ws.to(1).come("hello").find(3).admin("world")
	fmt.Println(ws.a,ws.b,ws.c,ws.d)


}
