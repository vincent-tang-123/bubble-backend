package controller

 import (
	 "bubble/dao"
	 "bubble/models"
	 "github.com/gin-gonic/gin"
	 "net/http"
 )

/*
业务分层
url --> controller --> serve --> model
请求来了  控制器         业务逻辑    模型层的增删改查
*/


func CreateATodo(c *gin.Context){

	// 前端页面填写代办事项 点击提交 会发送请求到这里
	// 1 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2 存入数据库
	err := models.CreateATodo(&todo)
	// 3 返回响应
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}


func GetTodoList(c *gin.Context){
	todoList, err := models.GetAllTodo()
	if err = dao.DB.Find(&todoList).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error_msg": err.Error(),
		})
	}
	c.JSON(http.StatusOK, todoList)
}


func UpdateATodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	todo, err := models.GetATodo(id)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}


func DeleteATodo(c *gin.Context) {
	id, _:= c.Params.Get("id")
	if err := models.DeleteATodo(id); err!=nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
