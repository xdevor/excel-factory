package main

import (
	"echoException/excel-faker/router"
)

func main() {
	router := router.SetUp()
	router.Run()
}
