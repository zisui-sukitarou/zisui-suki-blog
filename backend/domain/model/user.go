package model

import (
	"time"
)

type User struct {
	Id          UserId
	Name        UserName
	Email       UserEmail
	Password    UserPassword
	DisplayName UserDisplayName
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

/* constructor */
type CommandNewUser struct {
	Name        UserName
	Email       UserEmail
	Password    UserPassword
	DisplayName UserDisplayName
}

func NewUser(command CommandNewUser) (*User, error) {
	return &User{
		Name:        command.Name,
		Email:       command.Email,
		Password:    command.Password,
		DisplayName: command.DisplayName,
	}, nil
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

/* update email */
func (u *User) UpdateDisplayName(displayName UserDisplayName) error {
	u.DisplayName = displayName
	return nil
}
