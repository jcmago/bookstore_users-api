package app

import (
	"github.com/jcmago/bookstore_users-api/controllers/ping"
	"github.com/jcmago/bookstore_users-api/controllers/users"
)

func mapUrls() {
	// Health Check
	router.GET("/ping", ping.Ping)

	// Users
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	//router.GET("/users/search", controllers.SearchUser)
}
