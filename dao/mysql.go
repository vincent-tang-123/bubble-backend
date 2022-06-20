package dao

import "github.com/jinzhu/gorm"
import 	_ "github.com/jinzhu/gorm/dialects/mysql" // _ 下划线代表 只导入不引用


var (
	DB *gorm.DB
)


func InitMysql()(err error){
	dsn := "root:521314td@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn) // 不写冒号 使用全局的变量 DB
	if err != nil{
		return
	}
	return DB.DB().Ping() // ping 不通 返回错误
}

func Close(){
	DB.Close()
}