package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetHeroes(c *gin.Context) {
	name := c.GetHeader("name")
	password := c.GetHeader("password")
	// if there is no name and password
	if name != "" && password != "" {
		// if service.CheckAuth(name, password) {

		// }
		// c.JSON(200, data)
	}
	fmt.Println(password)

}

func GetHero(c *gin.Context) {
	name := c.GetHeader("name")
	password := c.GetHeader("password")
	if name != "" && password != "" {
		// if service.CheckAuth(name, password) {

		// }
		// c.JSON(200, data)
	}
}
