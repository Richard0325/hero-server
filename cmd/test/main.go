package main

import (
	"hero-server/service"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.Debug("CheckAuth")
	service.CheckAuth("hahow", "rocks")
	// log.Debug("TakeAllHeroes")
	// service.TakeAllHeroes()
	// log.Debug("TakeAllHeroesWithProfiles")
	// service.TakeAllHeroesWithProfiles()
	// log.Debug("TakeSingleHero")
	// service.TakeSingleHero("3")
	// log.Debug("TakeSingleHeroWithProfile")
	// service.TakeSingleHeroWithProfile("3")
}
