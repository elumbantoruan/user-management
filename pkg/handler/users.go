package handler

import (
	"edison-takehome/pkg/crypto"
	"edison-takehome/pkg/model"
	"edison-takehome/pkg/repository"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.Repository
}

func NewUserHandler(repo repository.Repository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (u *UserHandler) GetUserList(c *gin.Context) {
	token := c.GetHeader("x-authentication-token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing authentication token"})
		return
	}

	_, err := crypto.ValidateEmailInToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userAccts, err := u.repo.GetUserAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userEmails []model.UserEmail
	for _, acct := range userAccts {
		userEmails = append(userEmails, model.UserEmail{
			Email: acct.Email,
			User: model.User{
				FirstName: acct.FirstName,
				LastName:  acct.LastName,
			},
		})
	}
	userList := model.UserList{Users: userEmails}

	c.JSON(http.StatusOK, userList)
}

func (u *UserHandler) UpdateUser(c *gin.Context) {
	token := c.GetHeader("x-authentication-token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing authentication token"})
		return
	}

	email, err := crypto.ValidateEmailInToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.validate(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userEmail := model.UserEmail{
		Email: email,
		User: model.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}

	_, err = u.repo.UpdateUserAccount(userEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (u *UserHandler) validate(user model.User) error {
	var errorMsg string
	if user.FirstName == "" {
		errorMsg += "firstname is required"
	}
	if user.LastName == "" {
		errorMsg += "lastname is required"
	}
	if len(errorMsg) > 0 {
		return errors.New(errorMsg)
	}
	return nil
}
