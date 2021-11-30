package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TaskAllocate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tasks": []interface{}{"B07FKM7SLJ","B004QOAS8A"}})
}
