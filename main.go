package main

import (
	"fmt"
	"gitlab.com/yjagdale/siem-data-producer/app"
)

func init() {
	fmt.Println("Starting Data Producer")
}

func main() {
	fmt.Println("Started Service")
	app.StartApplication()
}
