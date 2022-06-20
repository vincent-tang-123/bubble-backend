package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
)

func main(){
	// 连接数据库
	err := dao.InitMysql()
	if err != nil {
		fmt.Printf("%#v\n", err)
		panic(err)
	}
	// 程序退出 关闭数据库//
	defer dao.Close()

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	// 启动路由
	r:= routers.SetupRouter()
	r.Run(":9000")
}