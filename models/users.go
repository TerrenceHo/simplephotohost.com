package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// ErrNotFound is returned when resource cannot be found in database
	ErrNotFound = errors.New("models: resource not found")

	// ErrInvalidID is returned when an invalid ID is provided to a method like
	// Delete
	ErrInvalidID = errors.New("models: ID provided not found")
)

const userPwPepper = "$2a$06yxCG8px5KYNhqK/ZgBxHKuK7bIZ3q1X3qL6oKUyQc6Bk9kUoKabsK"

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	// defer db.Close()
	db.LogMode(true)
	return &UserService{
		db: db,
	}, nil
}

type UserService struct {
	db *gorm.DB
}

// Look up user by the ID provided
// 1 - user, nil
// 2 - nil, ErrNotFound
// 3 - nil, OtherError
func (us *UserService) ByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// Looks up a user with the given email address and returns that user.
func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// first will query using the provided gorm.db and it will get the first item
// returned and place it into dst.  If nothing is found in the query, it will
// return ErrNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// Create the provided user and backfill data
// like the ID, CreatedAt, nd UpdateAt field
func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper) // converting to byte slice
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return us.db.Create(user).Error
}

// Update takes in a pointer to a user and updates the database
// where the user is saved of the data in the passed in user
func (us *UserService) Update(user *User) error {
	return us.db.Save(user).Error
}

// Delete will delete the user with the provided id
func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

// Closes db connnection
func (us *UserService) Close() error {
	return us.db.Close()
}

// Drops user table and rebuilds it
func (us *UserService) DestructiveReset() error {
	if err := us.db.DropTableIfExists(&User{}).Error; err != nil {
		return err
	}
	return us.AutoMigrate()
}

// Attempts to automatically migrate the users table
func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"` //ignore this in database
	PasswordHash string `gorm:"not null"`
}
