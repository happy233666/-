package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	v := r.Group("/users")
	{
		v.POST("/register", Register)
		v.POST("/login", Login)
		v.PUT("/changePassword", ChangePassword)
	}
	m := r.Group("/warehouse")
	{
		m.POST("/warehouse", NewWarehouse)
		m.PUT("/warehouse", ChangeWarehouse)
		m.DELETE("/warehouse", DeleteIteam)
		m.PUT("/warehouse", ChangeIteam)
	}
	r.Run()
}
