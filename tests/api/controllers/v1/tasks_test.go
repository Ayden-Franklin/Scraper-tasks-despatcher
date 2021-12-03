package v1_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"tasks-dispatcher/internal/apiserver/api/controllers/v1"
	"testing"
)

func TestTaskAllocate(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	api.TaskAllocate(c)

	assert.Equal(t, 200, w.Code)

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, got)
}