package main

import (
	"os"
	"todoapp/http/handler"
	"todoapp/http/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "up"})
	})

	r.POST("/login", handler.LoginHandler)
	api := r.Group("api")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/me", func(ctx *gin.Context) {
			userId := ctx.MustGet("userId")
			userType := ctx.MustGet("user_type")
			ctx.JSON(200, gin.H{
				"userId": userId,
				"type":   userType,
			})
		})
		api.GET("/todos", handler.GetTodos)
		api.GET("/todobyid/:id", handler.GetTodoById)
		api.GET("/todolists", handler.GetTodoLists)
		api.PATCH("/complete/:id", handler.CompleteTodo)
		api.PATCH("/changemessage", handler.ChangeMessage)
		api.POST("/addtodo", handler.AddTodo)
		api.DELETE("/deletetodo", handler.DeleteTodo)
		api.DELETE("/deletelist", handler.DeleteTodoList)
		api.POST("/addlist", handler.AddTodoList)

	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
