// ProAndCon project main.go
package main

import (
	"fmt"
	"time"
)

func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
		fmt.Println("produce value is:", i)
		//time.Sleep(10 * time.Second)
	}
}

func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println("consumer value is:", v)
	}
}

func main() {
	//fmt.Println("生产者消费者模拟")
	c := make(chan int, 1)
	cc := make(chan int, 1)

	go Consumer(cc)
	go Producer(cc)

	for i := 0; i < 10; i++ {
		go func(queue chan<- int,ii int) {
			//fmt.Println(ii)
			queue<-ii
			//c <- ii
			fmt.Println("################",ii)

			//}
		}(c,i)
	}
	for i:=0;i<10;i++{

	<-c

	}
	fmt.Println("end")

	time.Sleep(15 * time.Second)

}
