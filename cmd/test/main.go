package main

import (
	"fmt"
	"hero-server/model"
	"hero-server/service"

	log "github.com/sirupsen/logrus"
)

/*
* This file is a playground to testify functionalities while developing
 */
func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	service.Init(model.DaoTypeMock, model.ModeBroken)
	_, err := service.TakeSingleHero("aabbcc")
	fmt.Println(err.Error())
}
