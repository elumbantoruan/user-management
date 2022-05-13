package handler

import (
	"edison-takehome/pkg/crypto"
	"edison-takehome/pkg/model"
	"edison-takehome/pkg/repository"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	repo repository.Repository
}

func NewLogin(repo repository.Repository) *Login {
	return &Login{
		repo: repo,
	}
}

func (l *Login) UserLogin(c *gin.Context) {
	var userLogin model.EmailPassword
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := l.validate(userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	acct, err := l.repo.GetUserAccount(userLogin.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if acct.Password != crypto.Hash(userLogin.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	userToken, err := crypto.CreateResponseWithToken(userLogin.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userToken)
}

func (l *Login) validate(userLogin model.EmailPassword) error {
	var errorMsg string
	if userLogin.Email == "" {
		errorMsg += "email is required"
	}
	if userLogin.Password == "" {
		errorMsg += "password is required"
	}
	if len(errorMsg) > 0 {
		return errors.New(errorMsg)
	}
	return nil
}
