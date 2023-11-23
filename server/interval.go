package server

import (
	log "github.com/sirupsen/logrus"
	"github.com/robfig/cron"
)

func SetInterval(f func()){
	c := cron.New()
	err := c.AddFunc(Config.RequestInterval, f)
	if err != nil {
		log.Fatalln(err)
	}
	c.Start()
}