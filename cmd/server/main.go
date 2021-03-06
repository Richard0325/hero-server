package main

import (
	"hero-server/controller"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.WarnLevel)
}

func main() {
	r := gin.Default()
	r.GET("/heroes", controller.GetHeroes)
	r.GET("/heroes/:heroId", controller.GetSingleHero)
	r.Run(":8080")
}
