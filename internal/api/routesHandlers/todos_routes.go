package routeshandlers

import (
	"github.com/GuillaumeDeconinck/todos-go/internal/api/dao"
	"github.com/GuillaumeDeconinck/todos-go/pkg/models"
	"github.com/GuillaumeDeconinck/todos-go/pkg/tools"
	"github.com/gin-gonic/gin"
)

func listTodos(c *gin.Context) {
	var ownerUuid = c.Query("ownerUuid")
	var todos, _ = dao.ListTodos(&ownerUuid)
	c.JSON(200, todos)
}

func getTodo(c *gin.Context) {
	var uuidToGet = c.Param("uuidToGet")
	// var ownerUuid = c.Query("ownerUuid")
	todo, err := dao.GetTodo(&uuidToGet)

	if err != nil {
		var httpError = ConvertToHttpError(err)
		HandleError(c, httpError)
		return
	}

	c.JSON(200, todo)
}

func createTodo(c *gin.Context) {
	var todoToCreate models.Todo
	err := c.Bind(&todoToCreate)
	if err != nil {
		tools.SugaredLogger.Errorf("Error while deserializing body: %s", err)
		c.Status(400)
		return
	}
	err = dao.CreateTodo(&todoToCreate)
	if err != nil {
		tools.SugaredLogger.Errorf("Error while creating todo: %s", err)
		c.Status(500)
		return
	}

	c.Status(201)
}

func updateTodo(c *gin.Context) {
	var todoToUpdate models.Todo
	err := c.Bind(&todoToUpdate)
	if err != nil {
		tools.SugaredLogger.Errorf("Error while deserializing body: %s", err)
		c.Status(400)
		return
	}
	err = dao.UpdateTodo(&todoToUpdate)
	if err != nil {
		tools.SugaredLogger.Errorf("Error while updating todo: %s", err)
		c.Status(500)
		return
	}

	c.Status(204)
}

func deleteTodo(c *gin.Context) {
	var uuidToDelete = c.Param("uuidToDelete")
	dao.DeleteTodo(&uuidToDelete)
	c.Status(204)
}

func AddPingRoutesHandlers(r *gin.Engine) {
	r.GET("/todos", listTodos)
	r.GET("/todos/:uuidToGet", getTodo)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:uuidToUpdate", updateTodo)
	r.DELETE("/todos/:uuidToDelete", deleteTodo)
}
