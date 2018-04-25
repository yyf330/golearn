package main

import (
    "fmt"
    "time"
)

const MYFILE = "logfile.log"

func main() {
    c := time.Tick(10 * time.Second)
    for _ = range c {
        fmt.Println("dd")
    }
}
