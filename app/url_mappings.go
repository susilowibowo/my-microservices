package app

import "github.com/susilowibowo/my-microservices/controllers/ping"

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
