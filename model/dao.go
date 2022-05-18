package model

import (
	"hero-server/data"
)

type Dao interface {
	CallListHeroes() (data.Heroes, error)
	CallSingleHero(id string) (*data.Hero, error)
	CallProfileOfHero(id string) (*data.Profile, error)
	CallAuthenticate(name string, password string) (bool, error)
}

type DaoType int

var DaoTypeHahow DaoType = 0
var DaoTypeMock DaoType = 1

func InitDao(t DaoType) Dao {
	if t == DaoTypeHahow {
		return HahowDao{}
	} else {
		return MockDao{}
	}
}
