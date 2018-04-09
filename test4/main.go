package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	//var pre, next, sum int
	var sum int
	pre := 0
	next := 1
	count := -1
	fmt.Println("-------1---------",count)
	return func() int {
		count++
		if count < 2 {
			return count
		}
		fmt.Println("---------2-------",count)
		sum = pre + next
		pre = next
		next = sum
		fmt.Println("---------next-------",next)
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
