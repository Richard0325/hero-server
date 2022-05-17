package dao

import (
	"hero-server/data"
	"strconv"
	"testing"
)

func TestCallListHeroes(t *testing.T) {
	got, err := CallListHeroes()
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
	} else if err != data.ErrHahowServer1000 {
		t.Error("dao: CallListHeroes error")
	}
}

func TestCallSingleHero(t *testing.T) {
	for i := 1; i < 5; i++ {
		id := strconv.Itoa(i)
		got, err := CallSingleHero(id)
		if err == nil {
			wantId := id
			if got.Id != wantId {
				t.Error("dao: CallSingleHero error")
			}
		} else if err != data.ErrHahowServer1000 {
			t.Error("dao: CallSingleHero error")
		}
	}
	got, err := CallSingleHero("fkmdkfl")
	if got != nil && err != data.ErrIdNotFound {
		t.Error("dao: CallSingleHero error")
	}
	got, err = CallSingleHero("9999")
	if got != nil && err != data.ErrIdNotFound {
		t.Error("dao: CallSingleHero error")
	}
}

func TestCallProfileOfHero(t *testing.T) {
	got, err := CallProfileOfHero("2")
	want := data.Profile{
		Str: 8,
		Int: 2,
		Agi: 5,
		Luk: 9,
	}
	if err == nil {
		if got.Str != want.Str || got.Int != want.Int || got.Agi != want.Agi || got.Luk != want.Luk {
			t.Error("dao: CallProfileOfHero error")
		}
	} else if err != data.ErrHahowServer1000 {
		t.Error("dao: CallProfileOfHero error")
	}

	_, err = CallProfileOfHero("-10")
	if err != data.ErrIdNotFound {
		t.Error("dao: CallProfileOfHero error")
	}
}

func TestCallAuthenticate(t *testing.T) {
	got, err := CallAuthenticate("hahow", "rocks")
	if err == nil && got == true {
		t.Log("dao: CallAuthenticate pass")
	} else if err == data.ErrHahowServer1000 {
		t.Log("dao: CallAuthenticate pass")
	} else {
		t.Error("dao: CallAuthenticate error")
	}
	_, err = CallAuthenticate("whatever", "it is")
	if err == data.ErrNotAuthed {
		t.Log("dao: CallAuthenticate pass")
	} else {
		t.Error("dao: CallAuthenticate error")
	}

}
