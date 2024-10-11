package main

import (
	"github.com/parvejmia9/healthythakun/backend/router"
	"github.com/parvejmia9/healthythakun/database"
)

func main() {
	database.ConnectToMongoDB()
	defer func() {
		database.DisconnectToMongoDB()
	}()
	r := router.InitRouter()
	err := router.HandleReq(r)
	if err != nil {
		return
	}
}
