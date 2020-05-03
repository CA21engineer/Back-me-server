package main

import (
	"ca-zoooom/infrastructure"
	"fmt"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
	}
	infrastructure.Router.Run(fmt.Sprintf(":%s", port))
}
