package service

import (
	"hero-server/data"
	"hero-server/model"
	"hero-server/tools"

	log "github.com/sirupsen/logrus"
)

var dao model.Dao

/*
* To decide what dao to use, Real Hahow server or Mock server?
 */
func Init(t model.DaoType, m model.MockMode) {
	if t == model.DaoTypeMock {
		dao = model.InitMock(m)
	} else {
		dao = model.InitHahow()
	}
}

/*
* To Check if the user is authorized.
* Return true only when receiving "OK"
 */
func CheckAuth(name string, password string) (bool, error) {
	isAuthed, err := dao.CallAuthenticate(name, password)
	//While hahow API error happens, trys to call it again and again till reach 10 times
	cnt := 0
	for err == data.ErrHahowServer1000 {
		isAuthed, err = dao.CallAuthenticate(name, password)
		cnt++
		if cnt == 10 {
			log.Warn("Hahow Authenticate API error")
			return false, data.ErrRequestTimeout
		}
	}
	log.Debug("CheckAuth return: ", isAuthed)
	return isAuthed, err
}

/*
* To get all the heroes from Hahow
 */
func TakeAllHeroes() (data.Heroes, error) {
	ret, err := dao.CallListHeroes()
	//While hahow API error happens, trys to call it again and again till reach 10 times
	cnt := 0
	for err == data.ErrHahowServer1000 {
		ret, err = dao.CallListHeroes()
		cnt++
		if cnt == 10 {
			log.Warn("Hahow ListHeroes API error")
			return nil, data.ErrRequestTimeout
		}
	}
	log.Debug("TakeAllHeros return: ", tools.PrettyPrint(ret))
	return ret, err
}

/*
* To get all the heroes and put profile for each hero
 */
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

/*
* To get single hero
 */
func TakeSingleHero(id string) (*data.Hero, error) {
	ret, err := dao.CallSingleHero(id)
	cnt := 0
	//While hahow API error happens, trys to call it again and again till reach 10 times
	for err == data.ErrHahowServer1000 {
		ret, err = dao.CallSingleHero(id)
		cnt++
		if cnt == 10 {
			log.Warn("Hahow Single Hero API error")
			return nil, data.ErrRequestTimeout
		}
	}
	log.Debug("TakeSingleHero return: ", tools.PrettyPrint(ret))
	return ret, err
}

/*
* To get single hero and put profile to him
 */

func TakeSingleHeroWithProfile(id string) (*data.AuthHero, error) {
	hero, err := TakeSingleHero(id)
	if err != nil {
		return nil, err
	}
	profile, err := dao.CallProfileOfHero(hero.Id)
	if err != nil {
		log.Warn("service: dao.CallProfileOfHero error", err)
		return nil, err
	}
	ret := data.AuthHero{
		Id:      hero.Id,
		Name:    hero.Name,
		Image:   hero.Image,
		Profile: profile,
	}
	log.Debug("TakeSingleHeroWithProfile return: ", tools.PrettyPrint(ret))
	return &ret, nil
}
