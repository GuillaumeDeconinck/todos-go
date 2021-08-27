package routes

import (
	"net/http"

	"github.com/GuillaumeDeconinck/todos-go/internal/api/dao"
	"github.com/GuillaumeDeconinck/todos-go/pkg/models"
	"github.com/GuillaumeDeconinck/todos-go/pkg/tools"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func listTodos(c *gin.Context) {
	var ownerUuid = c.Query("ownerUuid")
	var todos, _ = dao.ListTodos(&ownerUuid)
	c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	var uuidToGet = c.Param("uuidToGet")
	// Todo: ownerUuid should be mandatory (it should/will be provided by the auth JWT)
	// var ownerUuid = c.Query("ownerUuid")
	todo, err := dao.GetTodo(&uuidToGet)

	if err != nil {
		var httpError = ConvertToHttpError(err)
		HandleError(c, httpError)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func createTodo(c *gin.Context) {
	var todoToCreate models.Todo
	err := c.Bind(&todoToCreate)
	if err != nil {
		tools.SugaredLogger.Errorf("Error while deserializing body: %s", err)
		c.Status(http.StatusBadRequest)
		return
	}

	err = validate.Struct(todoToCreate)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = dao.CreateTodo(&todoToCreate)
	if err != nil {
		var httpError = ConvertToHttpError(err)
		HandleError(c, httpError)
		return
	}

	c.Status(http.StatusCreated)
}

func updateTodo(c *gin.Context) {
	var todoToUpdate models.Todo
	err := c.Bind(&todoToUpdate)
	if err != nil {
		tools.SugaredLogger.Errorf("Error while deserializing body: %s", err)
		c.Status(http.StatusBadRequest)
		return
	}

	err = validate.Struct(todoToUpdate)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = dao.UpdateTodo(&todoToUpdate)
	if err != nil {
		var httpError = ConvertToHttpError(err)
		HandleError(c, httpError)
		return
	}

	c.Status(http.StatusNoContent)
}

func deleteTodo(c *gin.Context) {
	var uuidToDelete = c.Param("uuidToDelete")
	dao.DeleteTodo(&uuidToDelete)
	c.Status(http.StatusNoContent)
}

func AddPingRoutesHandlers(r *gin.Engine) {
	r.GET("/todos", listTodos)
	r.GET("/todos/:uuidToGet", getTodo)
	r.POST("/todos", createTodo)
	r.PUT("/todos/:uuidToUpdate", updateTodo)
	r.DELETE("/todos/:uuidToDelete", deleteTodo)
}
