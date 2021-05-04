package services

import (
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
)

var (
	HealthService healthServiceInterface = &health{}
)

type health struct{}

func (h *health) HealthCheck() *http_utils.Response {
	return http_utils.NewOkResponse("Pong")
}

type healthServiceInterface interface {
	HealthCheck() *http_utils.Response
}
