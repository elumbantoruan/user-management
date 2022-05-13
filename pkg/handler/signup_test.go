package handler

import (
	"bytes"
	"edison-takehome/pkg/model"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignup_UserSignup(t *testing.T) {
	w := httptest.NewRecorder()
	userAccount := model.UserAccount{
		FirstName: "Fname",
		LastName:  "Lname",
		Email:     "someone@gmail.com",
		Password:  "pwd",
	}
	b, err := json.Marshal(&userAccount)
	assert.NoError(t, err)

	r := io.NopCloser(bytes.NewBuffer(b))

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Body: r,
	}

	signup := NewSignup(&MockRepository{})
	signup.UserSignup(ctx)

	assert.Equal(t, http.StatusCreated, w.Code)

}
