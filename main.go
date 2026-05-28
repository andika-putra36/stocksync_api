package main

import (
	"stocksync_api/internal/config"
)

func main() {
	router := config.InitializeEverything()
	router.Run(":8888")
}
