package main

import "hello/internal/router"

func main() {
	r := router.Setup()
	r.Run(":8080")
}
