package main

import (
	routers "tasks-dispatcher/internal/apiserver/api/routers"
)

func main() {
	//setup routes
	router := routers.SetupRouter()

	// running
	router.Run()
}