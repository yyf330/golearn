package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func main()  {
    buf,err := ioutil.ReadFile("key.txt")
    if err != nil {
        log.Fatal(err)
    }
    statistic_times := make(map[string]int)
    words_length := strings.Fields(string(buf)) //在切片存放

    for _ , word := range words_length {
        _ , ok :=statistic_times[word]  //判断key是否存在，这个word是字符串。
        if ok{
            statistic_times[word] = statistic_times[word] + 1
        }else {
            statistic_times[word] = 1
        }
    }
    for word,counts := range statistic_times {
        fmt.Println(word,counts)
    }
}
