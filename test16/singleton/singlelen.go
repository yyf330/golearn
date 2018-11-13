package singleton

import (
    "fmt"
    "sync"
)

type Singleton interface {
    SaySomething()
}

type singleton struct {
    text string
}

var oneSingleton Singleton

var once sync.Once


func NewSingleton(tt string) Singleton {
    once.Do(func() {
        oneSingleton = &singleton{
            text: tt,
        }
    })

    return oneSingleton
}

func (this *singleton) SaySomething() {
    fmt.Println(this.text)
}
