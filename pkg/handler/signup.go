package handler

import (
	"net/http"

	"edison-takehome/pkg/crypto"
	"edison-takehome/pkg/model"
	"edison-takehome/pkg/repository"

	"github.com/gin-gonic/gin"
)

type Signup struct {
	repo repository.Repository
}

func NewSignup(repo repository.Repository) *Signup {
	return &Signup{
		repo: repo,
	}
}

func (s *Signup) UserSignup(c *gin.Context) {
	var userSignup model.UserAccount
	if err := c.ShouldBindJSON(&userSignup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.repo.InsertUserAccount(userSignup)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userToken, err := crypto.CreateResponseWithToken(userSignup.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userToken)
}
