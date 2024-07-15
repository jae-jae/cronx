# Cronx

A cross-platform lightweight task scheduler.

## Introduction

Cronx is a cross-platform, lightweight task scheduler designed to simplify the management of scheduled tasks. It allows you to define and execute tasks with ease using simple configuration files. Whether you're automating scripts or setting up regular maintenance tasks, Cronx provides a flexible and efficient solution.

## Installation

You can install Cronx using one of the following methods:

### Using `go install`

Ensure you have Golang installed, then run the following command:

```shell
go install github.com/jae-jae/cronx@latest
```

### Downloading from GitHub Releases
Alternatively, you can download the pre-built binary from the [releases page](https://github.com/jae-jae/cronx/releases)  on GitHub. Choose the appropriate version for your operating system, download the binary, and add it to your system's PATH.


## Usage

Run a cron task:
```shell
cronx
```

Specify the configuration file path: 
```shell
cronx -c cornx.yaml
```

Run tasks once, typically for testing purposes:


```shell
cronx run task1 task2
# or
cronx -c cornx.yaml run task1 task2
```

## Configuration
The configuration file `cronx.yaml` is used to define tasks and their schedules. By default, Cronx looks for this file in the current directory.

### Basic Configuration

Below is an example of a basic configuration:


```yaml
# scheduled task list
tasks:
  task1:
    schedule: "@every 1s"
    commands:
      - echo "hello world1!"
```

### Complete Configuration

Here is an example of a more complete configuration:


```yaml
# global Settings
settings:
  timezone: "UTC"

# set global environment variables.
env:
  LANG: en_US.UTF-8

# scheduled task list
tasks:
  task1:
    schedule: "@every 1s"
    commands:
      - echo "hello world1!"

  task2:
    schedule: "@every 2s"
    dir: /tmp
    commands:
      - echo "hello world2!"
      - echo $key1 > 1.txt
      - ls -la /tmp/
    env:
      key1: val1
```

### Configuration Instructions
The following table describes the fields used in the `cronx.yaml` configuration file:

| Field                   | Default                     | Description                                   |
|-------------------------|-----------------------------|-----------------------------------------------|
| settings                |                             | Global settings                               |
| settings.timezone       | (current server time zone)  | Set the cron job time zone                    |
| env                     |                             | Global environment variables                  |
| tasks                   |                             | Scheduled task list                           |
| tasks.[taskID]          |                             | Unique task ID                                |
| tasks.[taskID].schedule |                             | Cron expression                               |
| tasks.[taskID].dir      | (current working directory) | Directory for command execution               |
| tasks.[taskID].commands |                             | List of commands                              |
| tasks.[taskID].env      |                             | Environment variables for the current command |

## Cron Expressions
For detailed documentation, refer to the [robfig/cron documentation](https://pkg.go.dev/github.com/robfig/cron) .

A cron expression consists of six space-separated fields:

```shell
* * * * * *
| | | | | |
| | | | | +-- Seconds (0 - 59)
| | | | +---- Minutes (0 - 59)
| | | +------ Hours (0 - 23)
| | +-------- Day of the month (1 - 31)
| +---------- Month (1 - 12) (or JAN-DEC)
+------------ Day of the week (0 - 6) (or SUN-SAT)
```

### Predefined Schedules

You can use predefined schedules instead of a cron expression:

| Entry                  | Description                                | Equivalent To |
|------------------------|--------------------------------------------|---------------|
| @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *   |
| @monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *   |
| @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0   |
| @daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *   |
| @hourly                | Run once an hour, beginning of hour        | 0 0 * * * *   |

### Intervals

You can also schedule jobs to execute at fixed intervals:


```shell
@every <duration>
```
where "duration" is a string accepted by `time.ParseDuration` (http://golang.org/pkg/time/#ParseDuration). 

For example, `@every 1h30m10s` schedules a job to run every 1 hour, 30 minutes, and 10 seconds.

Note: The interval does not account for the job runtime. If a job takes 3 minutes to run and is scheduled to run every 5 minutes, there will be 2 minutes of idle time between runs.

## Related Projects

- [robfig/cron](https://github.com/robfig/cron)