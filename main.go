package main

import (
	"ca-zoooom/infrastructure"
)

func main() {
	infrastructure.Router.Run(":8084")
}
