package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User is a moderator user that has access to the control panel
type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Username    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"password"`
	Email       string             `bson:"email" json:"email"`
	Fullname    string             `bson:"fullname" json:"fullname"`
	Permissions Permissions        `bson:"permissions" json:"permissions"`
}

// NewUser is a user factory that generates a user with no permissions and a hashed password
func NewUser(
	username string,
	password string,
	email string,
	fullname string,
) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	permissions := NewPermissions(false, false, false, false)
	return &User{
		ID:          primitive.NewObjectID(),
		Username:    username,
		Password:    string(hashedPassword),
		Email:       email,
		Fullname:    fullname,
		Permissions: *permissions,
	}, nil
}

// GetID returns the hex of the user's id
func (u *User) GetID() string {
	return u.ID.Hex()
}

// GetObjectID returns the object id of the user
func (u *User) GetObjectID() primitive.ObjectID {
	return u.ID
}

// GetUsername returns the user's username
func (u *User) GetUsername() string {
	return u.Username
}

// SetUsername changes the user's username
func (u *User) SetUsername(newUsername string) error {
	if newUsername == "" {
		return errors.New("Username can't be empty")
	}
	u.Username = newUsername
	return nil
}

// GetEmail returns the user's email
func (u *User) GetEmail() string {
	return u.Email
}

// SetEmail sets a new email for the user object
func (u *User) SetEmail(newEmail string) error {
	u.Email = newEmail
	return nil
}

// GetFullname returns the user's fullname
func (u *User) GetFullname() string {
	return u.Fullname
}

// SetFullname sets a new fullname for the user object
func (u *User) SetFullname(newFullname string) {
	u.Fullname = newFullname
}

// SetPermissions sets the permissions of the user
func (u *User) SetPermissions(users bool, tickets bool, events bool, sponsors bool) {
	u.Permissions.Users = users
	u.Permissions.Tickets = tickets
	u.Permissions.Events = events
	u.Permissions.Sponsors = sponsors
}

// ChangePassword sets a new encrypted password for the user
func (u *User) ChangePassword(newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
