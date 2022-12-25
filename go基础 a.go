package main

//结构体必须是大写字母开头的成员才会被JSON处理到，小写字母开头的成员不会有影响。

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type User struct {
	Username string
	Nickname string
	Sex      uint8
	Birthday time.Time
}

func main() {
	u := User{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
