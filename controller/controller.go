package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetHeroes(c *gin.Context) {
	name := c.GetHeader("Name")
	password := c.GetHeader("Password")
	if name != "" && password != "" {

		// isAuthed, _ := service.CheckAuth(name, password)
		// if isAuthed {

		// }
		// c.JSON(200, data)
	}
	fmt.Println(password)

}

func GetHero(c *gin.Context) {
	name := c.GetHeader("Name")
	password := c.GetHeader("Password")
	if name != "" && password != "" {
		// if service.CheckAuth(name, password) {

		// }
		// c.JSON(200, data)
	}
}
