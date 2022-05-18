package controller

import (
	"hero-server/data"
	"hero-server/model"
	"hero-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	service.Init(model.DaoTypeHahow, model.ModeNone)
}

func GetHeroes(c *gin.Context) {
	name := c.GetHeader("Name")
	password := c.GetHeader("Password")
	if name != "" && password != "" {
		isAuthed, _ := service.CheckAuth(name, password)
		if isAuthed {
			authHeroes, err := service.TakeAllHeroesWithProfiles()
			if err != nil {
				c.JSON(http.StatusInternalServerError, nil)
				return
			}
			c.JSON(http.StatusOK, authHeroes)
			return
		}
	}
	heroes, err := service.TakeAllHeroes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, heroes)
}

func GetHero(c *gin.Context) {
	id := c.Param("id")
	name := c.GetHeader("Name")
	password := c.GetHeader("Password")
	if name != "" && password != "" {
		isAuthed, _ := service.CheckAuth(name, password)
		if isAuthed {
			authHero, err := service.TakeSingleHeroWithProfile(id)
			if err == data.ErrIdNotFound {
				c.JSON(http.StatusNotFound, "NotFound")
				return
			} else if err != nil {
				c.JSON(http.StatusInternalServerError, nil)
				return
			}
			c.JSON(http.StatusOK, authHero)
			return
		}
	}
	hero, err := service.TakeSingleHero(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "NotFound")
		return
	}
	c.JSON(http.StatusOK, hero)
}
