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

// FindAllUsers in database
func (u *UserDB) FindAllUsers() (*[]model.User, error) {
	users := []model.User{}

	if err := u.cl.Model(&model.User{}).Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}
