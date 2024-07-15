package cron

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Settings *Settings         `yaml:"settings"`
	Env      map[string]string `yaml:"env"`
	Tasks    map[string]*Task  `yaml:"tasks"`
}

type Settings struct {
	Timezone string `yaml:"timezone"`
}

type Task struct {
	Schedule string            `yaml:"schedule"`
	Commands []string          `yaml:"commands"`
	Env      map[string]string `yaml:"env,omitempty"`
	Dir      string            `yaml:"dir,omitempty"`
}

func (c *Config) MergeEnv(env map[string]string) []string {
	envs := os.Environ()
	if c.Env != nil {
		for k, v := range c.Env {
			envs = append(envs, k+"="+v)
		}
	}

	if env != nil {
		for k, v := range env {
			envs = append(envs, k+"="+v)
		}
	}

	return envs
}

func LoadConfig(path string) (*Config, error) {
	if path == "" {
		path = "./cronx.yaml"
	}

	log.Infof("load config from %s", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
