package routeshandlers

import (
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err *HttpError) {
	c.JSON(err.StatusCode, &err)
}
