package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main(){
	//TODO 连接数据库
	//初始化路由
	db, err := gorm.Open("mysql", "root:991030@(localhost)/golang?charset=utf8mb4&parseTime=True&loc=Local")
	if err!=nil{
		panic(err)
	}
	defer db.Close()
	router :=gin.Default()
	r1 := router.Group("/openapi/db")
	{
		r1.GET("/list/:tablename")
		r1.GET("/one/:tablename")
		r1.POST("/new")
	}
	_ = router.Run(":8000")
}
