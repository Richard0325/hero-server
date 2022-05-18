package service

import (
	"hero-server/data"
	"hero-server/model"
	"testing"
)

func TestCheckAuth(t *testing.T) {
	got, err := CheckAuth("hahow", "rocks")
	if err == nil && got != true {
		t.Error("service.CheckAuth error")
	} else if got == false && err == data.ErrRequestTimeout {
		t.Log("service.CheckAuth pass")
	}
	got, err = CheckAuth("whatever", "it is")
	if got != false || err == nil {
		t.Log(got)
		t.Log(err.Error())
		t.Error("service.CheckAuth error")
	}
}

func TestTakeAllHeroes(t *testing.T) {
	got, err := TakeAllHeroes()
	want := []data.Hero{
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
	}

	if err == nil {
		for i := range got {
			if got[i].Id == want[i].Id && got[i].Name == want[i].Name && got[i].Image == want[i].Image {
				continue
			} else {
				t.Error("dao: CallListHeroes error")
			}
		}
	}

}

func TestTakeAllHeroesWithProfiles(t *testing.T) {
	got, err := TakeAllHeroesWithProfiles()
	want := []data.AuthHero{
		{
			Id:    "1",
			Name:  "Daredevil",
			Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
			Profile: &data.Profile{
				Str: 2,
				Int: 7,
				Agi: 9,
				Luk: 7,
			},
		},
		{
			Id:    "2",
			Name:  "Thor",
			Image: "http://x.annihil.us/u/prod/marvel/i/mg/5/a0/537bc7036ab02/standard_xlarge.jpg",
			Profile: &data.Profile{
				Str: 8,
				Int: 2,
				Agi: 5,
				Luk: 9,
			},
		},
		{
			Id:    "3",
			Name:  "Iron Man",
			Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/a0/55b6a25e654e6/standard_xlarge.jpg",
			Profile: &data.Profile{
				Str: 6,
				Int: 9,
				Agi: 6,
				Luk: 9,
			},
		},
		{
			Id:    "4",
			Name:  "Hulk",
			Image: "http://i.annihil.us/u/prod/marvel/i/mg/5/a0/538615ca33ab0/standard_xlarge.jpg",
			Profile: &data.Profile{
				Str: 10,
				Int: 1,
				Agi: 4,
				Luk: 2,
			},
		},
	}
	if err == nil {
		for i := range got {
			if got[i].Id != want[i].Id || got[i].Name != want[i].Name || got[i].Image != want[i].Image {
				t.Error("value error")
			}
			gotProf := got[i].Profile
			wantProf := want[i].Profile
			if gotProf.Str != wantProf.Str || gotProf.Int != wantProf.Int || gotProf.Agi != wantProf.Agi || gotProf.Luk != wantProf.Luk {
				t.Error("value error")
			}
		}
	}
}

func TestTakeSingleHero(t *testing.T) {
	got, err := TakeSingleHero("1")
	if err == nil {
		want := data.AuthHero{
			Id:    "1",
			Name:  "Daredevil",
			Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
		}
		if got.Id != want.Id || got.Name != want.Name || got.Image != want.Image {
			t.Error("value error")
		}
	}
}

func TestTakeSingleHeroWithProfile(t *testing.T) {
	got, err := TakeSingleHeroWithProfile("1")
	if err == nil {
		want := data.AuthHero{
			Id:    "1",
			Name:  "Daredevil",
			Image: "http://i.annihil.us/u/prod/marvel/i/mg/6/90/537ba6d49472b/standard_xlarge.jpg",
			Profile: &data.Profile{
				Str: 2,
				Int: 7,
				Agi: 9,
				Luk: 7,
			},
		}
		if got.Id != want.Id || got.Name != want.Name || got.Image != want.Image {
			t.Error("value error")
		}
		gotProf := got.Profile
		wantProf := want.Profile
		if gotProf.Str != wantProf.Str || gotProf.Int != wantProf.Int || gotProf.Agi != wantProf.Agi || gotProf.Luk != wantProf.Luk {
			t.Error("value error")
		}
	}
}

func TestM(t *testing.T) {
	t.Log("Test Case 1: When Hahow API works normally")
	Init(model.DaoTypeMock, model.ModeNormal)
	TestCheckAuth(t)
	TestTakeAllHeroes(t)
	TestTakeAllHeroesWithProfiles(t)
	TestTakeSingleHero(t)
	TestTakeSingleHeroWithProfile(t)
	t.Log("Test Case 2: When Hahow API return error all the time")
	Init(model.DaoTypeMock, model.ModeBroken)
	TestCheckAuth(t)
	TestTakeAllHeroes(t)
	TestTakeAllHeroesWithProfiles(t)
	TestTakeSingleHero(t)
	TestTakeSingleHeroWithProfile(t)
	t.Log("Test Case 2: When unexpected error happend in Dao")
	Init(model.DaoTypeMock, model.ModeUnexpected)
	TestCheckAuth(t)
	TestTakeAllHeroes(t)
	TestTakeAllHeroesWithProfiles(t)
	TestTakeSingleHero(t)
	TestTakeSingleHeroWithProfile(t)
}
