package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func InitMysql()(err error){
	dsn := "root:521314td@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) // 不写 冒号 使用全局的变量

}


func main(){
	// 连接数据库

	// sql : create database bubble

	r := gin.Default()
	//r.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"error_msg": "",
	//		"data": "data",
	//	})
	//
	//})
	//r.GET("/v1/todo", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": 200,
	//		"error_msg": "",
	//		"data": "data",
	//	})
	//
	//})

	// v1
	v1Group := r.Group("/v1")
	{

	}

	// 待办事项
	// 添加
	v1Group.POST("/todo", func(c *gin.Context) {

	})
	// 查看所有的代办是事项
	v1Group.GET("/todo", func(c *gin.Context) {
		
	})
	// 查看某一个代办事项
	v1Group.GET("/todo/:id", func(c *gin.Context) {
		
	})
	// 修改
	v1Group.PUT("/todo/:id", func(c *gin.Context) {

	})
	// 删除
	v1Group.DELETE("/todo/:id", func(c *gin.Context) {

	})

	r.Run(":9000")
}