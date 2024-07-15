package cron

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var configFile string

type CMD struct {
	rootCmd *cobra.Command
}

func NewCMD() *CMD {
	return &CMD{}
}

func (c *CMD) Run() error {
	c.initRootCmd()
	c.runTaskCmd()
	return c.rootCmd.Execute()
}

func (c *CMD) initRootCmd() {
	c.rootCmd = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			config, err := LoadConfig(configFile)
			if err != nil {
				log.Fatalf("load config failed: %v", err)
			}
			cancel, err := c.runCron(config)
			if err != nil {
				log.Fatalf("run cron failed: %v", err)
			}
			defer cancel()

			s := make(chan os.Signal, 1)
			signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)
			sig := <-s
			log.Infof("cron closing down (signal: %v)", sig)
		},
	}
	c.rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is ./cronx.yaml)")
}

func (c *CMD) runTaskCmd() {
	c.rootCmd.AddCommand(&cobra.Command{
		Use:   "run [tasks]",
		Short: "Run tasks once, usually for testing purposes.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			config, err := LoadConfig(configFile)
			if err != nil {
				log.Fatalf("load config failed: %v", err)
			}
			for _, task := range args {
				executor := TaskExecutor{
					Config: config,
					TaskID: task,
				}
				err := executor.Exec()
				if err != nil {
					log.Errorf("task %s exec failed: %v", task, err)
				}
			}
		},
	})
}

func (c *CMD) runCron(config *Config) (func(), error) {
	ce, err := NewCronExecutor(config)
	if err != nil {
		return nil, err
	}

	for taskID, task := range config.Tasks {
		_, err := ce.AddFunc(task.Schedule, func() {
			executor := TaskExecutor{
				Config: config,
				TaskID: taskID,
			}
			err := executor.Exec()
			if err != nil {
				log.Errorf("task %s exec failed: %v", taskID, err)
			}
		})
		if err != nil {
			return nil, err
		}
	}

	ce.Start()

	return func() {
		ce.Stop()
	}, nil
}
