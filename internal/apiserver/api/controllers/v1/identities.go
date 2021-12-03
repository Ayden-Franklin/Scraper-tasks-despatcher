package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	util "tasks-dispatcher/pkg/util"
	database "tasks-dispatcher/internal/apiserver/durable"
)


func IdentityRetrieve(c *gin.Context) {
	id := util.Generate()
	database.SaveIdentity(id)
	c.JSON(http.StatusOK, gin.H{"id": id})
}