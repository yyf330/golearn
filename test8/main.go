package main

import (
    "bufio"
    "os"
    "fmt"

    "strconv"
)

var   (
    s string
    line string
)
func main()  {
    f := bufio.NewReader(os.Stdin)
    num := []int{100,200,300,400,500,600,700,800}
    fmt.Printf("现有一些数字：·\033[32;1m%v\033[0m·\n",num)
    for {
        fmt.Print("请您想要反转下标的起始的位置>")
        line,_ = f.ReadString('\n')
        if len(line) == 1 {
            continue  //过滤掉空格；
        }
        fmt.Sscan(line,&s)
        if s == "stop" {
            break //定义停止程序的键值;
        }
        index,err := strconv.Atoi(s)
        if err != nil || index <1 || index > 8{
            fmt.Println("对不起，您必须输入一个大于0小于8的数字")
	    continue
        }
        num1 := num[:index]
        num2 := num[index:]
        fmt.Printf("·\033[31;1m%v\033[0m·\n",num1,num2)
        i := index-1
        for {

            num2=append(num2, num1[i])
            i = i - 1
            if len(num2)>=8 {
                break
            }
        }
        fmt.Printf("反转后的内容是·\033[31;1m%v\033[0m·\n",num2)
    }
}

