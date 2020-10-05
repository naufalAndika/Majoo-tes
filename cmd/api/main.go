package main

import (
	"github.com/naufalAndika/Majoo-tes/cmd/api/server"
	"github.com/naufalAndika/Majoo-tes/internal/platform/mysql"
)

func main() {
	e := server.New()

	db, err := mysql.New("root:rahasia123@(localhost:3306)/tes_majoo?charset=utf8mb4&parseTime=true")
	if err != nil {
		panic(err)
	}

	mysql.Seed(db)

	server.Start(e)
}
