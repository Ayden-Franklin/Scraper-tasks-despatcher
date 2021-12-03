package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	api_v1 "tasks-dispatcher/internal/apiserver/api/controllers/v1"
)

func  SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})

		v1.GET("/", func(c *gin.Context){
			c.String(http.StatusOK, "Hello")
		})
		v1.GET("/identity", api_v1.IdentityRetrieve)
		v1.GET("/allocate", api_v1.TaskAllocate)
	}
	return r
}
