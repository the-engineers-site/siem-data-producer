package main

import (
	"fmt"
	"siem-data-producer/app"
	"siem-data-producer/utils/logger_utils"
)

func init() {
	logger_utils.LoggerUtils.InitLogger()
}

func main() {
	fmt.Println("Started Service")
	app.StartApplication()
}
