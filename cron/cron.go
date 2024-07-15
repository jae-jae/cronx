package cron

import (
	"fmt"
	cronExecutor "github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewCronExecutor(config *Config) (*cronExecutor.Cron, error) {
	var opts []cronExecutor.Option
	if config.Settings != nil && config.Settings.Timezone != "" {
		loc, err := time.LoadLocation(config.Settings.Timezone)
		if err != nil {
			return nil, fmt.Errorf("load timezone %s failed: %v", config.Settings.Timezone, err)
		}
		log.Infof("set cron timezone to %s", config.Settings.Timezone)
		opts = append(opts, cronExecutor.WithLocation(loc))
	}
	return cronExecutor.New(opts...), nil
}
