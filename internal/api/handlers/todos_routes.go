package handlers

import (
	"github.com/GuillaumeDeconinck/todos-go/internal/api/dao"
	"github.com/GuillaumeDeconinck/todos-go/pkg/models"
	"github.com/GuillaumeDeconinck/todos-go/pkg/tools"
	"github.com/gin-gonic/gin"
)

// type TodoCreateDTO struct {
// 	Uuid        *string `gorm:"primaryKey"`
// 	OwnerUuid   *string `gorm:"index"`
// 	State       *string
// 	Title       *string
// 	Description *string
// }

func listTodos(c *gin.Context) {
	var ownerUuid = c.Query("ownerUuid")
	var todos, _ = dao.ListTodos(&ownerUuid)
	c.JSON(200, todos)
}

func getTodo(c *gin.Context) {

}

func createTodo(c *gin.Context) {
	var todoCreateDTO models.Todo
	err := c.Bind(&todoCreateDTO)
	if err != nil {
		tools.SugaredLogger.Errorf("Error while deserializing body: %s", err)
		c.Status(400)
		return
	}
	dao.CreateTodo(&todoCreateDTO)

	c.Status(201)
}

func updateTodo(c *gin.Context) {

}

func deleteTodo(c *gin.Context) {
	var uuidToDelete = c.Param("uuidToDelete")
	dao.DeleteTodo(&uuidToDelete)
	c.Status(204)
}

func AddPingRoutesHandlers(r *gin.Engine) {
	r.GET("/todos", listTodos)
	r.POST("/todos", createTodo)
	r.DELETE("/todos/:uuidToDelete", deleteTodo)
}
