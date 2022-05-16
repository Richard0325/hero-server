package data

type Profile struct {
	Str int `json:"str"`
	Int int `json:"int"`
	Agi int `json:"agi"`
	Luk int `json:"luk"`
}

type Hero struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Profile *Profile `json:"profile"`
}

type Heroes []*Hero
