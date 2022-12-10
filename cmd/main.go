package main

import (
	"gateway/internal/router"
)

func main() {
	r := router.Routers()
	r.Run(":8080")

}
