package cron

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Run() {
	initLog()

	cmd := NewCMD()
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func initLog() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
}
