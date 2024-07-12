package cron

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mvdan.cc/sh/v3/expand"
	"mvdan.cc/sh/v3/interp"
	"mvdan.cc/sh/v3/syntax"
	"os"
)

type TaskExecutor struct {
	Config *Config
	TaskID string
}

func (t *TaskExecutor) Exec() error {
	taskConfig, ok := t.Config.Tasks[t.TaskID]
	if !ok {
		return errors.New("task not found")
	}

	envs := t.Config.MergeEnv(taskConfig.Env)
	for _, cmd := range taskConfig.Commands {
		executor := CMDExecutor{
			Cmd:  cmd,
			Envs: envs,
			Dir:  taskConfig.Dir,
		}

		err := executor.Exec()
		if err != nil {
			return fmt.Errorf("command [%s] exec failed: %v", cmd, err)
		}
	}

	return nil
}

type CMDExecutor struct {
	Cmd  string
	Envs []string
	Dir  string
}

func (c *CMDExecutor) Exec() error {
	// Parse the script
	parser := syntax.NewParser()
	file, err := parser.Parse(bytes.NewReader([]byte(c.Cmd)), "")
	if err != nil {
		return err
	}

	// Prepare the interpreter
	runner, err := interp.New(
		interp.StdIO(nil, os.Stdout, os.Stderr),
		interp.Env(expand.ListEnviron(c.Envs...)),
		interp.Dir(c.Dir),
	)
	if err != nil {
		return err
	}

	// Execute the script
	if err := runner.Run(context.Background(), file); err != nil {
		return err
	}

	return nil
}
