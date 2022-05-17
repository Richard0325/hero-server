package tools

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func PrettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		log.Panic("PrettyPrint Error", err.Error())
	}
	return string(s)
}
