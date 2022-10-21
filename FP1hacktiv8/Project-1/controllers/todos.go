package controllers

import (
	"FP1hacktiv8/models"
	"fmt"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All Todos
// @Summary Get All Data Todos
// @Description Get All Todos
// @Accept json
// @Produce json
// @Router /todos [get]
func GetAllTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetAllTodos())
}

// Get All Todos By ID
// @Summary Get Data Todos
// @Description Get Todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todos"
// @Param id path int true "ID Todo"
// @Router /todos/{id} [get]
func GetAllTodosByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	todo, err := models.GetTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data Not Found",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

// Create Todo
// @Summary Post New Data Todos
// @Description Post New Data Todos
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessCreate
// @Param todo body models.Todo true "Todos"
// @Param id path int true "ID Todo"
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data can not created",
		})
		return
	}

	models.AddNewTodo(&newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// Update Todo
// @Summary Edit Data Todos
// @Description Edit Data Todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Update"
// @Success 200 {object} response.SuccessUpdate
// @Param id path int true "ID Todo"
// @Router /todos/update/{id} [put]
func UpdateTodo(c *gin.Context) {
	var updateTodo models.Todo

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	if err := c.BindJSON(&updateTodo); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data can not created",
		})
		return
	}
	todo, errUpdate := models.UpdateTodoById(id, &updateTodo)
	if errUpdate != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data Not Found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"msg": "Data" + todo.ToList + "Has Been Update",
	})
}

// Delete Todo
// @Summary Delete Data Todos
// @Description Delete Data Todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Delete"
// @Success 200 {object} response.SuccessDelete
// @Param id path int true "ID Todo"
// @Router /todos/delete/{id} [delete]
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}

	todo, err := models.DeleteTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "Data Not Found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"msg": "Data" + todo.ToList + "Has Been Deleted",
	})
}
