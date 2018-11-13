package main

import (

	"fmt"
	"io/ioutil"
	"log"
	//"path/filepath"
	"strings"
	//"os"
)

func main() {
	//buf,err := ioutil.ReadFile("/root/work/src/test/test10/key.txt")
	//if err != nil {
	//    log.Fatal(err)
	//}
	//statistic_times := make(map[string]int)
	//words_length := strings.Fields(string(buf)) //在切片存放
	//
	//for _ , word := range words_length {
	//    _ , ok :=statistic_times[word]  //判断key是否存在，这个word是字符串。
	//    if ok{
	//        statistic_times[word] = statistic_times[word] + 1
	//    }else {
	//        statistic_times[word] = 1
	//    }
	//}
	//for word,counts := range statistic_times {
	//    fmt.Println(word,counts)
	//}
	//path := filepath.Join("config1")
	//
	////文件类型需要进行过滤
	//fmt.Println(os.Getwd())
	buf, err := ioutil.ReadFile("config")
	if err != nil {
		log.Fatal(err)
	}
	content := string(buf)
	fmt.Println(content)
	if strings.Contains(content,"192.168.1.153"){
		fmt.Println("####################33")
	}
	//替换
	newContent := strings.Replace(content, "192.168.1.153", "192.168.1.154", -1)
	fmt.Println(newContent)


}
