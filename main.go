package main

import (
	"fmt"
	"gitlab.com/yjagdale/siem-data-producer/app"
	"gitlab.com/yjagdale/siem-data-producer/utils"
)

func init() {
	utils.LoggerUtils.InitLogger()
}

func main() {
	fmt.Println("Started Service")
	app.StartApplication()
}
