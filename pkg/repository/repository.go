package repository

import "edison-takehome/pkg/model"

type Repository interface {
	InsertUserAccount(acct model.UserAccount) error
	UpdateUserAccount(userEmail model.UserEmail) (*model.UserAccount, error)
	GetUserAccount(email string) (*model.UserAccount, error)
	GetUserAccounts() ([]model.UserAccount, error)
	Close() error
}
