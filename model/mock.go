package model

import "hero-server/data"

type MockMode int

const ModeNone MockMode = -1
const ModeNormal MockMode = 0     // return normal data
const ModeBroken MockMode = 1     // hahow server code 1000 instead of normal data
const ModeUnexpected MockMode = 2 // unexpected error (like http request library errors and so on)

type MockDao struct {
	Mode MockMode
}

func InitMock(m MockMode) MockDao {
	return MockDao{
		Mode: m,
	}
}

func (m MockDao) CallListHeroes() (data.Heroes, error) {
	if m.Mode == ModeNormal {
		return data.Heroes{
			{
				Id:    "1",
				Name:  "Daredevil",
				Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
			},
			{
				Id:    "2",
				Name:  "Thor",
				Image: "http://x.annihil.us/u/prod/marvel/i/mg/5/a0/537bc7036ab02/standard_xlarge.jpg",
			},
			{
				Id:    "3",
				Name:  "Iron Man",
				Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/a0/55b6a25e654e6/standard_xlarge.jpg",
			},
			{
				Id:    "4",
				Name:  "Hulk",
				Image: "http://i.annihil.us/u/prod/marvel/i/mg/5/a0/538615ca33ab0/standard_xlarge.jpg",
			},
		}, nil
	} else if m.Mode == ModeBroken {
		return nil, data.ErrHahowServer1000
	}
	return nil, data.ErrUnknown
}

func (m MockDao) CallSingleHero(id string) (*data.Hero, error) {
	if m.Mode == ModeNormal {
		switch id {
		case "1":
			return &data.Hero{
				Id:    "1",
				Name:  "Daredevil",
				Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
			}, nil
		case "2":
			return &data.Hero{
				Id:    "1",
				Name:  "Daredevil",
				Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
			}, nil
		case "3":
			return &data.Hero{
				Id:    "1",
				Name:  "Daredevil",
				Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
			}, nil
		case "4":
			return &data.Hero{
				Id:    "1",
				Name:  "Daredevil",
				Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
			}, nil
		default:
			return nil, data.ErrIdNotFound
		}
	} else if m.Mode == ModeBroken {
		switch id {
		case "1", "2", "3", "4":
			return nil, data.ErrHahowServer1000
		default:
			return nil, data.ErrIdNotFound
		}
	}
	return nil, data.ErrUnknown
}

func (m MockDao) CallProfileOfHero(id string) (*data.Profile, error) {
	if m.Mode == ModeNormal {
		switch id {
		case "1":
			return &data.Profile{
				Str: 2,
				Int: 7,
				Agi: 9,
				Luk: 7,
			}, nil
		case "2":
			return &data.Profile{
				Str: 8,
				Int: 2,
				Agi: 5,
				Luk: 9,
			}, nil
		case "3":
			return &data.Profile{
				Str: 6,
				Int: 9,
				Agi: 6,
				Luk: 9,
			}, nil
		case "4":
			return &data.Profile{
				Str: 10,
				Int: 1,
				Agi: 4,
				Luk: 2,
			}, nil
		default:
			return nil, data.ErrIdNotFound
		}
	} else if m.Mode == ModeBroken {
		switch id {
		case "1", "2", "3", "4":
			return nil, data.ErrHahowServer1000
		default:
			return nil, data.ErrIdNotFound
		}
	}
	return nil, data.ErrUnknown
}

func (m MockDao) CallAuthenticate(name string, password string) (bool, error) {
	if m.Mode == ModeNormal {
		if name == "hahow" && password == "rocks" {
			return true, nil
		} else {
			return false, data.ErrNotAuthed
		}
	} else if m.Mode == ModeBroken {
		if name == "hahow" && password == "rocks" {
			return false, data.ErrHahowServer1000
		} else {
			return false, data.ErrNotAuthed
		}
	}
	return false, data.ErrUnknown
}
