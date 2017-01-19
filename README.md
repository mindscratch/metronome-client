# metronome-client

Metronome client library written in Go inspired by [chronos-client](https://github.com/yieldbot/chronos-client/).
Some code was also copied/reused where practical.

This library targets the `v1` API.

## Run Examples

Example programs exist for most API calls. To run example, such as `print_jobs`, do the following:

```
$ go run examples/jobs/print_jobs.go
[
  {
    "id": "prod.example.app",
    "description": "Example Application",
    "labels": {
      "location": "olympus",
      "owner": "zeus"
    },
    "run": {
      "args": null,
      "artifacts": [],
      "cmd": "env | sort",
      "cpus": 0.1,
      "disk": 128,
      "docker": {
        "image": ""
      },
      "env": {
        "CONNECT": "direct",
        "MON": "test"
      },
      "maxLaunchDelay": 3600,
      "mem": 32,
      "placement": {
        "constraints": []
      },
      "user": "root",
      "restart": {
        "policy": "NEVER",
        "activeDeadlineSeconds": 0
      },
      "volumes": []
    }
  }
```