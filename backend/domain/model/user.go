package model

import (
	"time"
)

type User struct {
	UserId         UserId
	Name           UserName
	Email          UserEmail
	HashedPassword UserHashedPassword
	Icon           UserIcon
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

/* constructor */
func NewUser(
	userId UserId,
	name UserName,
	email UserEmail,
	hashedPassword UserHashedPassword,
	icon UserIcon,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		UserId:         userId,
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
		Icon:           icon,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}
}

/* update user_id */
func (u *User) UpdateUserId(userId UserId) {
	u.UserId = userId
	u.UpdatedAt = time.Now()
}

/* update name */
func (u *User) UpdateName(name UserName) {
	u.Name = name
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

/* update password */
func (u *User) UpdatePassword(hashedPassword UserHashedPassword) {
	u.HashedPassword = hashedPassword
	u.UpdatedAt = time.Now()
}
