package main

import (
	"fmt"
	"time"
)
//令牌桶算法的原理是系统会以一个恒定的速度往桶里放入令牌，而如果请求需要被处理，则需要先从桶里获取一个令牌，当桶里没有令牌可取时，则拒绝服务。从原理上看，令牌桶算法和漏桶算法是相反的，一个“进水”，一个是“漏水”。
func logs(args ...interface{}) {
	fmt.Println(args...)
}

func tokenBucket(limit int, rate int) chan struct{} {
	tb := make(chan struct{}, limit)
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	for i := 0; i < limit; i++ {
		tb <- struct{}{}
	}

	go func() {
		for {
			for i := 0; i < rate; i++ {
				tb <- struct{}{}
			}
			<-ticker.C
		}
	}()
	return tb
}

func popToken(bucket chan struct{}, n int) {
	for i := 0; i < n; i++ {
		<-bucket
	}
}

func testTokenBucket() {
	rate := 10
	tb := tokenBucket(20, rate)

	dataLen := 100
	for i := 0; i <= dataLen; i += rate {
		popToken(tb, rate)
		logs(i)
	}
}

func main() {
	testTokenBucket()
}
