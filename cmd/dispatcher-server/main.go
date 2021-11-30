package main

import (
	routers "tasks-dispatcher/internal/apiserver/api/controllers"
)

func main() {
	//setup routes
	router := routers.SetupRouter()

	// running
	router.Run()
}