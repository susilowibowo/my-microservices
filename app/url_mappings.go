package app

import (
	"github.com/susilowibowo/my-microservices/controllers/ping"
	"github.com/susilowibowo/my-microservices/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
