package services

import (
	"gitlab.com/yjagdale/siem-data-producer/utils/http_utils"
)

var (
	HealthService healthServiceInterface = &health{}
)

type health struct{}

type healthServiceInterface interface {
	Check() *http_utils.ResponseEntity
}

func (h *health) Check() *http_utils.ResponseEntity {
	return http_utils.NewOkResponse("Pong")
}
