package model

import (
	"time"
)

type User struct {
	UserId    UserId
	Name      UserName
	Email     UserEmail
	Password  UserPassword
	CreatedAt time.Time
	UpdatedAt time.Time
}

/* constructor */
type CommandNewUser struct {
	UserId   UserId
	Name     UserName
	Email    UserEmail
	Password UserPassword
}

func NewUser(command CommandNewUser) (*User, error) {
	return &User{
		UserId:   command.UserId,
		Name:     command.Name,
		Email:    command.Email,
		Password: command.Password,
	}, nil
}

/* update user_id */
func (u *User) UpdateUserId(userId UserId) error {
	u.UserId = userId
	return nil
}

/* update name */
func (u *User) UpdateName(name UserName) error {
	u.Name = name
	return nil
}

/* update email */
func (u *User) UpdateEmail(email UserEmail) error {
	u.Email = email
	return nil
}

/* update password */
func (u *User) UpdatePassword(password UserPassword) error {
	u.Password = password
	return nil
}
