package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_GetUserList(t *testing.T) {

	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{}

	user := NewUserHandler(&MockRepository{})
	user.GetUserList(ctx)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"missing authentication token"}`, string(w.Body.Bytes()))
}
