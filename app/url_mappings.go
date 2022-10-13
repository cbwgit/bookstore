package app

import (
	"bookstore/controllers/ping"
	"bookstore/controllers/users"
)

func mapUrls() {

	//router.GET("/", users.Test)
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
