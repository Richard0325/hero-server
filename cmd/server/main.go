package main

import (
	"hero-server/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/heroes", controller.GetHeroes)
	r.GET("/heroes/:id", controller.GetHero)
	r.Run(":8080")
}
