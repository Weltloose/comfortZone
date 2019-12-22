package main

import (
	"github.com/Weltloose/comfortZone/router"
)

func main() {
	server := router.GetServer()
	server.Run()
}
