package handler

import (
	"net/http"
	"strconv"
	"todoapp/domain/model"
	"todoapp/usecase"

	"github.com/gin-gonic/gin"
)

func GetTodoLists(ctx *gin.Context) {
	userType, exist := ctx.Get("user_type")
	if !exist {
		return
	}
	if userType == 2 {
		ctx.JSON(http.StatusOK, usecase.AllTodoLists())
		return
	}

	id, exist := ctx.Get("userId")
	if !exist {
		return
	}
	todos := usecase.GetTodoListsByUserId(id.(int))
	ctx.JSON(http.StatusOK, todos)
}
func GetTodos(ctx *gin.Context) {
	userType, exist := ctx.Get("user_type")
	if !exist {
		return
	}
	if userType == 2 {
		ctx.JSON(http.StatusOK, usecase.AllTodos())
		return
	}

	id, exist := ctx.Get("userId")
	if !exist {
		return
	}
	todos := usecase.GetTodosByUserId(id.(int))
	ctx.JSON(http.StatusOK, todos)
}
func GetAllTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, usecase.AllTodos())
}
func AddTodo(ctx *gin.Context) {
	var step model.TodoStep
	if err := ctx.ShouldBindJSON(&step); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usecase.AddTodo(&step)
}

type TodoListInput struct {
	Title string `json:"title"`
}

func AddTodoList(ctx *gin.Context) {
	var input TodoListInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz JSON"})
		return
	}

	var list model.TodoList
	list.Title = input.Title
	usecase.CreateTodoList(&list, ctx.MustGet("userId").(int))
}
func CompleteTodo(ctx *gin.Context) {
	todoIdStr := ctx.Param("id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Geçersiz ID"})
		return
	}

	usecase.Complete(todoId)
	ctx.Status(200)

}

type ChangeMessageInput struct {
	TodoId  int    `json:"todoId"`
	Message string `json:"message"`
}

func ChangeMessage(ctx *gin.Context) {
	var input ChangeMessageInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz JSON"})
		return
	}

	usecase.ChangeMessage(input.TodoId, input.Message)
	ctx.Status(200)
}

type DeleteTodoInput struct {
	TodoId int `json:"todoId"`
}

func DeleteTodo(ctx *gin.Context) {
	var input DeleteTodoInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz JSON"})
		return
	}

	usecase.DeleteTodo(input.TodoId)
	ctx.Status(200)
}

type DeleteTodoListInput struct {
	ID int `json:"id"`
}

func DeleteTodoList(ctx *gin.Context) {
	var input DeleteTodoListInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz JSON"})
		return
	}

	usecase.DeleteTodoList(input.ID)
	ctx.Status(200)
}
func GetTodoById(ctx *gin.Context) {

	todoIdStr := ctx.Param("id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Geçersiz ID"})
		return
	}
	ctx.JSON(http.StatusOK, usecase.GetTodoById(todoId))
	ctx.Status(200)
}
