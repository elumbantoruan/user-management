package handler

import (
	"bytes"
	"edison-takehome/pkg/crypto"
	"edison-takehome/pkg/model"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin_UserLogin(t *testing.T) {

	w := httptest.NewRecorder()
	emailPassword := model.EmailPassword{
		Email:    "someone@gmail.com",
		Password: "pwd",
	}
	b, err := json.Marshal(&emailPassword)
	assert.NoError(t, err)

	r := io.NopCloser(bytes.NewBuffer(b))

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Body: r,
	}

	login := NewLogin(&MockRepository{})
	login.UserLogin(ctx)
	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestLogin_UserLogin_MismatchPassword_StatusUnauthorized(t *testing.T) {

	w := httptest.NewRecorder()
	emailPassword := model.EmailPassword{
		Email:    "someone@gmail.com",
		Password: "pwdZZZZ",
	}
	b, err := json.Marshal(&emailPassword)
	assert.NoError(t, err)

	r := io.NopCloser(bytes.NewBuffer(b))

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Body: r,
	}

	login := NewLogin(&MockRepository{})
	login.UserLogin(ctx)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

}

type MockRepository struct{}

func (mr *MockRepository) InsertUserAccount(acct model.UserAccount) error {
	return nil
}

func (mr *MockRepository) UpdateUserAccount(userEmail model.UserEmail) (*model.UserAccount, error) {
	return nil, nil
}

func (mr *MockRepository) GetUserAccount(email string) (*model.UserAccount, error) {
	return &model.UserAccount{
		Email:     "someone@gmail.com",
		FirstName: "fname",
		LastName:  "lname",
		Password:  crypto.Hash("pwd"),
	}, nil
}

func (mr *MockRepository) GetUserAccounts() ([]model.UserAccount, error) {
	return nil, nil
}

func (mr *MockRepository) Close() error {
	return nil
}
