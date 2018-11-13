package main

import (
    singleton "test/test16/singleton"
)

func main() {
    mSingleton, nSingleton := singleton.NewSingleton("hel1lo"), singleton.NewSingleton("hi")
    mSingleton.SaySomething()
    nSingleton.SaySomething()
}


//----------------------- goroutine 测试 ------------------------
//func main() {
//    c := make(chan int)
//    go newObject("hello", c)
//    go newObject("hi", c)
//
//    <-c
//    <-c
//}
//
//func newObject(str string, c chan int) {
//    nSingleton := singleton.NewSingleton(str)
//    nSingleton.SaySomething()
//    c <- 1
//}
