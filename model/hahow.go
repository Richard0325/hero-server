package model

import (
	"bytes"
	"encoding/json"
	"hero-server/data"
	"hero-server/tools"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HahowDao struct{}

func InitHahow() HahowDao {
	return HahowDao{}
}

func (h HahowDao) CallListHeroes() (data.Heroes, error) {
	url := "https://hahow-recruit.herokuapp.com/heroes"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Panic(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()
	// only 200 is acceptable
	if response.StatusCode != http.StatusOK {
		log.Panic(err)
		return nil, data.ErrUnknown
	}
	body, _ := ioutil.ReadAll(response.Body)
	ret := data.Heroes{}
	json.Unmarshal(body, &ret)
	if ret[0].Id == "" {
		log.Error("ListHeroes Backend Error")
		return nil, data.ErrHahowServer1000
	}
	log.Trace("CallListHeroes return: ", tools.PrettyPrint(ret))
	return ret, nil
}

func (h HahowDao) CallSingleHero(id string) (*data.Hero, error) {
	url := "https://hahow-recruit.herokuapp.com/heroes/" + id
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Panic(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	// only 404 or 200 is acceptable
	if response.StatusCode == http.StatusNotFound {
		log.Trace("SingleHero Id Not Found")
		return nil, data.ErrIdNotFound
	} else if response.StatusCode != http.StatusOK {
		return nil, data.ErrUnknown
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	ret := data.Hero{}
	json.Unmarshal(body, &ret)
	if ret.Id == "" {
		log.Trace("SingleHero Backend Error")
		return nil, data.ErrHahowServer1000
	}
	log.Trace(tools.PrettyPrint(ret))
	return &ret, nil
}

func (h HahowDao) CallProfileOfHero(id string) (*data.Profile, error) {
	url := "https://hahow-recruit.herokuapp.com/heroes/" + id + "/profile"
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Panic(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	// only 404 or 200 is acceptable
	if response.StatusCode == http.StatusNotFound {
		log.Info("ProfileOfHero Id Not Found")
		return nil, data.ErrIdNotFound
	} else if response.StatusCode != http.StatusOK {
		return nil, data.ErrUnknown
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	ret := data.Profile{}
	json.Unmarshal(body, &ret)
	if ret.Str == 0 && ret.Int == 0 && ret.Agi == 0 && ret.Luk == 0 {
		log.Trace("ProfileOfHero Backend Error")
		return nil, data.ErrHahowServer1000
	}
	log.Trace("profile return: ", tools.PrettyPrint(ret))
	return &ret, nil
}

func (h HahowDao) CallAuthenticate(name string, password string) (bool, error) {
	url := "https://hahow-recruit.herokuapp.com/auth"
	values := map[string]string{"name": name, "password": password}
	jsonValue, _ := json.Marshal(values)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Panic(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	// only 401 or 200 is acceptable
	if response.StatusCode == http.StatusUnauthorized {
		log.Trace("Not Authorized")
		return false, data.ErrNotAuthed
	} else if response.StatusCode != http.StatusOK {
		log.Trace("Unkown Error")
		return false, data.ErrUnknown
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	content := string(body)
	if content != "OK" {
		log.Trace("Authenticate Backend Error")
		return false, data.ErrHahowServer1000
	}
	log.Trace("auth return: ", content)
	return true, nil
}
