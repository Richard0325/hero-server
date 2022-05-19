package controller

import (
	"hero-server/data"
	"hero-server/model"
	"hero-server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
* Init Dao to access real Hahow server
 */
func init() {
	service.Init(model.DaoTypeHahow, model.ModeNone)
}

/*
* Handle API /heroes
 */
func GetHeroes(c *gin.Context) {
	name := c.GetHeader("Name")
	password := c.GetHeader("Password")
	//if there are name and password, check if they are authorized.
	if name != "" && password != "" {
		isAuthed, _ := service.CheckAuth(name, password)
		if isAuthed {
			authHeroes, err := service.TakeAllHeroesWithProfiles()
			if err != nil {
				c.JSON(http.StatusInternalServerError, "Not able to acces data")
				return
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"heroes": authHeroes,
			})
			return
		}
	}
	//all people without being authorized successfully will get unauthenticated data
	heroes, err := service.TakeAllHeroes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"heroes": heroes,
	})
}

/*
* Handle API /heroes/:heroId
 */

func GetSingleHero(c *gin.Context) {
	heroId := c.Param("heroId")
	name := c.GetHeader("Name")
	password := c.GetHeader("Password")
	//if there are name and password, check if they are authorized.
	if name != "" && password != "" {
		isAuthed, _ := service.CheckAuth(name, password)
		if isAuthed {
			authHero, err := service.TakeSingleHeroWithProfile(heroId)
			if err == data.ErrIdNotFound {
				c.JSON(http.StatusNotFound, "NotFound")
				return
			} else if err != nil {
				c.JSON(http.StatusInternalServerError, "Not able to acces data")
				return
			}
			c.JSON(http.StatusOK, authHero)
			return
		}
	}
	//all people without being authorized successfully will get unauthenticated data
	hero, err := service.TakeSingleHero(heroId)
	if err != nil {
		c.JSON(http.StatusNotFound, "NotFound")
		return
	}
	c.JSON(http.StatusOK, hero)
}
