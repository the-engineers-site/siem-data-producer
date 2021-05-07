package main

import (
	"fmt"
	"gitlab.com/yjagdale/siem-data-producer/app"
	"gitlab.com/yjagdale/siem-data-producer/utils/logger_utils"
)

func init() {
	logger_utils.LoggerUtils.InitLogger()
}

func main() {
	fmt.Println("Started Service")
	app.StartApplication()
}
