package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tasks-dispatcher/internal/apiserver/redis"
)

func TaskAllocate(c *gin.Context) {
	asin := redis.PopItem()

	c.JSON(http.StatusOK, gin.H{"tasks": asin})
}
