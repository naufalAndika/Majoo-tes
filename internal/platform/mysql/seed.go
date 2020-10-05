package mysql

import (
	"log"

	model "github.com/naufalAndika/Majoo-tes/internal"

	"github.com/jinzhu/gorm"
)

var users = []model.User{
	model.User{
		Username: "naufalandika",
		Password: "nopalnopal",
		Fullname: "M Naufal Andika",
	},
	model.User{
		Username: "dhaffaandika",
		Password: "dhaffadhaffa",
		Fullname: "M Dhaffa Andika",
	},
}

// Seed database
func Seed(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&model.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&model.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&model.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
