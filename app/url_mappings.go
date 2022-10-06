package app
import (
"bookstore/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	
}