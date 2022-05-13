package model

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserEmail struct {
	Email string `json:"email"`
	User
}

type UserEmailPassword struct {
	UserEmail
	Password string `json:"password"`
}

type EmailPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserList struct {
	Users []UserEmail `json:"users"`
}

type UserAccount struct {
	Email     string `json:"email" pg:",unique"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}
