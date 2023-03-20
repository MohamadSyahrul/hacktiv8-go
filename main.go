package main

import "go-learning/router"

func main() {
	router.StartServer().Run(":3000")
}
