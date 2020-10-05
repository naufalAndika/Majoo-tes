package mysql

import (
	"github.com/jinzhu/gorm"
	model "github.com/naufalAndika/Majoo-tes/internal"
)

// UserDB model
type UserDB struct {
	cl *gorm.DB
}

// NewUserDB create new UserDB instance
func NewUserDB(c *gorm.DB) *UserDB {
	return &UserDB{c}
}

// CreateUser store new user
func (u *UserDB) CreateUser(user model.User) (*model.User, error) {
	if err := u.cl.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAllUsers in database
func (u *UserDB) FindAllUsers() (*[]model.User, error) {
	users := []model.User{}

	if err := u.cl.Model(&model.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

// FindByID find user by id
func (u *UserDB) FindByID(ID uint) (*model.User, error) {
	var user model.User

	if err := u.cl.First(&user, ID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteByID delete user by id
func (u *UserDB) DeleteByID(ID uint) error {
	if err := u.cl.Delete(&model.User{}, ID).Error; err != nil {
		return err
	}

	return nil
}

// UpdateByID update user by id
func (u *UserDB) UpdateByID(ID uint, userData model.User) (*model.User, error) {
	var user model.User

	if err := u.cl.First(&user, ID).Error; err != nil {
		return nil, err
	}

	user.Password = userData.Password
	user.Fullname = userData.Fullname
	if err := u.cl.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByCredentials find user by username and password
func (u *UserDB) FindByCredentials(username, password string) (*model.User, error) {
	var user model.User
	if err := u.cl.Where("username = ?", username).Where("password = ?", password).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
