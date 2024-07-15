# cronx
A cross-platform lightweight task scheduler

## Usage

Run cron task:
```shell
cronx
```

Run tasks once, usually for testing purposes:
```shell
cronx run task1 task2
```


## Config

The simplest configuration:
```yaml
# scheduled task list
tasks:
  task1:
    schedule: "@every 1s"
    commands:
      - echo "hello world1!"
```

Complete configuration:
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

Configuration instructions for `cronx.yaml`:

| Field                   | Default                         | Description                                       |
|-------------------------|---------------------------------|---------------------------------------------------|
| settings                |                                 | global Settings                                   |
| settings.timezone       | (current server time zone)      | set cron job time zone                            |
| env                     |                                 | set global environment variables                  |
| tasks                   |                                 | scheduled task list                               |
| tasks.[taskID]          |                                 | define a unique task ID                           |
| tasks.[taskID].schedule |                                 | cron expression                                   |
| tasks.[taskID].dir      | (folder where cronx is running) | set the folder for command execution              |
| tasks.[taskID].commands |                                 | command list                                      |
| tasks.[taskID].env      |                                 | set environment variables for the current command |

## Cron expression

> Reference documentation: https://pkg.go.dev/github.com/robfig/cron

A cron expression represents a set of times, using 6 space-separated fields.
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

**Predefined schedules:**

You may use one of several pre-defined schedules in place of a cron expression.

| Entry                  | Description                                | Equivalent To |
|------------------------|--------------------------------------------|---------------|
| @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *   |
| @monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *   |
| @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0   |
| @daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *   |
| @hourly                | Run once an hour, beginning of hour        | 0 0 * * * *   |

**Intervals:**

You may also schedule a job to execute at fixed intervals, starting at the time it's added or cron is run. This is supported by formatting the cron spec like this:

```shell
@every <duration>
```

where "duration" is a string accepted by time.ParseDuration (http://golang.org/pkg/time/#ParseDuration).

For example, `@every 1h30m10s` would indicate a schedule that activates after 1 hour, 30 minutes, 10 seconds, and then every interval after that.

Note: The interval does not take the job runtime into account. For example, if a job takes 3 minutes to run, and it is scheduled to run every 5 minutes, it will have only 2 minutes of idle time between each run.

## Related projects
- [robfig/cron](https://github.com/robfig/cron)