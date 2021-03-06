# metronome-client

[![GoDoc](https://godoc.org/github.com/mindscratch/metronome-client?status.png)](https://godoc.org/github.com/mindscratch/metronome-client)

[Metronome](https://github.com/dcos/metronome) client library written in Go inspired by [chronos-client](https://github.com/yieldbot/chronos-client/).

This library targets the `v1` API.

## Issues affecting this library

It's worth noting that I've opened a couple [issues](https://github.com/dcos/metronome/issues/created_by/mindscratch) that I've found with Metronome while implementing this library.

## Examples

Look at the programs in the `examples` directory. An application exists for each v1 API endpoint, they also demonstrate how to use this library.

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
