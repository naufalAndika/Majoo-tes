package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// New DB connection
func New(destination string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", destination)
	if err != nil {
		return nil, err
	}

	db = db.Set("gorm:table_options", "DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci ENGINE=InnoDB")
	db = db.Set("gorm:auto_preload", true)

	return db, nil
}
