package app

import "gitlab.com/yjagdale/siem-data-producer/controllers/health"

func mapUrls() {
	router.GET("/ping", health.Ping)
}
