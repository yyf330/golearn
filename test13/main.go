package main

import (
	"fmt"
	"encoding/json"
)



type Server struct {
	// ID 不会导出到JSON中
	ID int `json:"-"`

	// ServerName2 的值会进行二次JSON编码
	ServerName string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// 如果 ServerIP 为空，则不输出到JSON串中
	ServerIP string `json:"serverIP,omitempty"`
}

type Serverslice struct {
	Servers []Server
}

func main() {
	s := []Server {
		Server{ID: 3,	ServerName: `Go "1.0" `,	ServerName2: `Go "1.0" `,	ServerIP: "192.168.1.1",},
		Server{ID: 4,	ServerName: `Go "1.8" `,	ServerName2: `Go "1.8" `,	ServerIP: "192.168.1.12",},
	}


	b, _ := json.Marshal(s)
	//fmt.Println(b)
	//os.Stdout.Write(b)
	fmt.Println(string(b))
	json.Unmarshal(b, &s)
	fmt.Println("parse===")
	fmt.Println(s)
	fmt.Println(s[0].ServerName)
}

