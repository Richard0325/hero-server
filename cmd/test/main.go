package main

import (
	"hero-server/dao"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.TraceLevel)
}

func main() {
	dao.CallAuthenticate("hahow", "rocks")
}
