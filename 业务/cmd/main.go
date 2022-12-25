package main

import (
	"gin/a/业务/api"
	"gin/a/业务/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
