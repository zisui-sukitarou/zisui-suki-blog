package model

import (
	"time"
)

type User struct {
	UserId      UserId
	Name        UserName
	DisplayName UserDisplayName
	Email       UserEmail
	Password    UserHashedPassword
	Icon        UserIcon
	Status      UserStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

/* constructor */
func NewUser(
	userId UserId,
	name UserName,
	displayName UserDisplayName,
	email UserEmail,
	password UserHashedPassword,
	icon UserIcon,
	status UserStatus,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		UserId:      userId,
		Name:        name,
		DisplayName: displayName,
		Email:       email,
		Password:    password,
		Icon:        icon,
		Status:      status,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

/* update name */
func (u *User) UpdateName(name UserName) {
	u.Name = name
	u.UpdatedAt = time.Now()
}

/* update name */
func (u *User) UpdateDisplayName(name UserDisplayName) {
	u.DisplayName = name
	u.UpdatedAt = time.Now()
}

/* update email */
func (u *User) UpdateEmail(email UserEmail) {
	u.Email = email
	u.UpdatedAt = time.Now()
}

/* update icon url */
func (u *User) UpdateIcon(icon UserIcon) {
	u.Icon = icon
	u.UpdatedAt = time.Now()
}

/* update icon url */
func (u *User) UpdateStatus(status UserStatus) {
	u.Status = status
	u.UpdatedAt = time.Now()
}

/* upadte password */
func (u *User) UpdatePassword(password UserHashedPassword) {
	u.Password = password
	u.UpdatedAt = time.Now()
}