package main

import (
	"fmt"
)

func test_unnamed()(int) {
	var i int
	defer func() {
		i++
		fmt.Println("defer a:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer b :", i)
	}()
	return i
}
func test_named()(i int) {
	defer func() {
		i++
		fmt.Println("defer c:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer d :", i)
	}()
	return i
}

func main() {
	fmt.Println("return:", test_unnamed())
	fmt.Println("return:", test_named())
	defer func(){
		fmt.Println("112")
	}()
	defer func(){
		fmt.Println("1333312")

	}()
	fmt.Println("dlskdjflj")
}
