package main

import (
	"github.com/naufalAndika/Majoo-tes/cmd/api/server"
)

func main() {
	e := server.New()
	server.Start(e)
}
