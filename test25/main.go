package main

import (
	"fmt"
	"time"
)

func main(){
	var i int



	for i=0 ; i<10 ; i++{
		defer fmt.Println(i)
		go func(i int) {
			fmt.Println("###",i)

		}(i)
		fmt.Println("--------",i)
	}
	time.Sleep(time.Second*1)

	panic("test")

	time.Sleep(time.Second*1)
}


