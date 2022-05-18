package controller

import (
	"encoding/json"
	"hero-server/data"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetHeroes(t *testing.T) {
	// create a router for /heroes
	router := gin.Default()
	router.GET("/heroes", GetHeroes)
	t.Log("test without name and password")
	// generate a new http request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/heroes", nil)
	// serve request to router
	router.ServeHTTP(w, req)
	// only status 500 and 200 is acceptable
	if w.Code != 500 {
		if w.Code != 200 {
			t.Error("controller: GetHeroes respond code other than 200")
		}
		var respData struct {
			Heroes data.Heroes `json:"heroes"`
		}
		// fill content of body into the data structure
		json.Unmarshal(w.Body.Bytes(), &respData)
		// check if there are id, name and image for each hero
		for _, hero := range respData.Heroes {
			if hero.Id == "" || hero.Name == "" || hero.Image == "" {
				t.Error("controller: GetHeroes return unexpected data")
			}
		}
	}
	t.Log("test with correct name and password")
	// generate a new http request
	w = httptest.NewRecorder()
	req.Header.Set("Name", "hahow")
	req.Header.Set("Password", "rocks")
	router.ServeHTTP(w, req)
	if w.Code != 500 {
		if w.Code != 200 {
			t.Error("controller: GetHeroes respond code other than 200")
		}
		var respData struct {
			Heroes data.AuthHeroes `json:"heroes"`
		}
		json.Unmarshal(w.Body.Bytes(), &respData)
		// check if there are id, name, image, profile for each hero
		for _, hero := range respData.Heroes {
			if hero.Id == "" || hero.Name == "" || hero.Image == "" {
				t.Error("controller: GetHeroes return unexpected data")
			}
			p := hero.Profile
			if p.Str == 0 && p.Int == 0 && p.Agi == 0 && p.Luk == 0 {
				t.Error("controller: GetHeroes return unexpected data")
			}
		}
	}
	t.Log("test with wrong name and password")
	w = httptest.NewRecorder()
	req.Header.Set("Name", "whatever")
	req.Header.Set("Password", "it is")
	// serve request to router
	router.ServeHTTP(w, req)
	// only status 500 and 200 is acceptable
	if w.Code != 500 {
		if w.Code != 200 {
			t.Error("controller: GetHeroes respond code other than 200")
		}
		var respData struct {
			Heroes data.AuthHeroes `json:"heroes"`
		}
		// fill content of body into the data structure
		json.Unmarshal(w.Body.Bytes(), &respData)
		// check if there are id, name, image but no profile for each hero
		for _, hero := range respData.Heroes {
			if hero.Id == "" || hero.Name == "" || hero.Image == "" {
				t.Error("controller: GetHeroes return unexpected data")
			}
			p := hero.Profile
			if p != nil {
				t.Error("controller: GetHeroes return profiles to unauthenticated client")
			}
		}
	}

}

func TestGetSingleHero(t *testing.T) {
	router := gin.Default()
	router.GET("/heroes/:heroId", GetSingleHero)
	t.Log("test with existed id and without using name and password")
	// generate a new http request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/heroes/1", nil)
	// serve request to router
	router.ServeHTTP(w, req)
	// only status 500 and 200 is acceptable
	if w.Code != 500 {
		if w.Code != 200 {
			t.Error("controller: GetSingleHero respond wrong status when heroId is 1")
		}
		var respData data.Hero
		// fill content of body into the data structure
		json.Unmarshal(w.Body.Bytes(), &respData)
		// check if id is correct and if name and image are not empty
		if respData.Id != "1" || respData.Name == "" || respData.Image == "" {
			t.Error("controller: GetSingleHero return unexpected data")
		}
	}
	t.Log("test with inexistent id and without using name and password")
	// generate a new http request
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/heroes/aabbcc", nil)
	// serve request to router
	router.ServeHTTP(w, req)
	// if status is not 500, it must be 404
	if w.Code != 500 {
		if w.Code != 404 {
			t.Error("controller: GetSingleHero respond wrong status when heroId is aabbcc")
		}
	}
	t.Log("test with existed id and using correct name and password")
	// generate a new http request
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/heroes/3", nil)
	req.Header.Set("Name", "hahow")
	req.Header.Set("Password", "rocks")
	// serve request to router
	router.ServeHTTP(w, req)
	// only status 500 and 200 is acceptable
	if w.Code != 500 {
		if w.Code != 200 {
			t.Error("controller: GetSingleHero respond wrong status when heroId is 3")
		}
		var respData data.AuthHero
		// fill content of body into the data structure
		json.Unmarshal(w.Body.Bytes(), &respData)
		// check if id is correct and if name and image are not empty
		if respData.Id != "3" || respData.Name == "" || respData.Image == "" {
			t.Error("controller: GetSingleHero return unexpected data")
		}
		// check if attibutes in profile are correct
		p := respData.Profile
		if p.Str != 4 || p.Int != 11 || p.Agi != 6 || p.Luk != 9 {
			t.Error("controller: GetSingleHero return wrong data when authenticated")

		}
	}
	t.Log("test with existed id and using wrong name or password")
	// generate a new http request
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/heroes/3", nil)
	req.Header.Set("Name", "whatever")
	req.Header.Set("Password", "it is")
	// serve request to router
	router.ServeHTTP(w, req)
	// only status 500 and 200 is acceptable
	if w.Code != 500 {
		if w.Code != 200 {
			t.Error("controller: GetSingleHero respond wrong status when heroId is 3")
		}
		var respData data.AuthHero
		// fill content of body into the data structure
		json.Unmarshal(w.Body.Bytes(), &respData)
		// check if id is correct and if name and image are not empty
		if respData.Id != "3" || respData.Name == "" || respData.Image == "" {
			t.Error("controller: GetSingleHero return unexpected data")
		}
		// profile should be nil
		if respData.Profile != nil {
			t.Error("controller: GetSingleHero return profile to unauthenticated client")
		}
	}
}
