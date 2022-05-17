package service

import (
	"hero-server/dao"
	"hero-server/data"
	"hero-server/tools"

	log "github.com/sirupsen/logrus"
)

func CheckAuth(name string, password string) (bool, error) {
	isAuth, err := dao.CallAuthenticate(name, password)
	cnt := 0
	for err == data.ErrHahowServer1000 {
		isAuth, err = dao.CallAuthenticate(name, password)
		cnt++
		if cnt == 10 {
			log.Warn("Hahow Authenticate API error")
			return false, err
		}
	}
	log.Debug("CheckAuth return: ", isAuth)
	return isAuth, nil
}

func TakeAllHeroes() (data.Heroes, error) {
	ret, err := dao.CallListHeroes()
	cnt := 0
	for err == data.ErrHahowServer1000 {
		ret, err = dao.CallListHeroes()
		cnt++
		if cnt == 10 {
			log.Warn("Hahow ListHeroes API error")
			return nil, err
		}
	}
	log.Debug("TakeAllHeros return: ", tools.PrettyPrint(ret))
	return ret, nil
}

func TakeAllHeroesWithProfiles() (data.AuthHeroes, error) {
	heroes, err := TakeAllHeroes()
	if err != nil {
		return nil, err
	}
	ret := data.AuthHeroes{}
	for _, hero := range heroes {
		authHero := data.AuthHero{
			Id:    hero.Id,
			Name:  hero.Name,
			Image: hero.Image,
		}
		profile, err := dao.CallProfileOfHero(hero.Id)
		if err != nil {
			log.Warn("service: dao.CallProfileOfHero error", err)
			return nil, err
		}
		authHero.Profile = profile
		ret = append(ret, &authHero)
	}
	log.Debug("TakeAllHeroesWithProfiles return: ", tools.PrettyPrint(ret))
	return ret, nil
}

func TakeSingleHero() {

}

func TakeSingleHeroWithProfile() {

}
