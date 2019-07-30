# AWASE server

## It does...

AWASE server gives REST-like API for choosing event holding date.

## It does not...

- authenticate request
- limit rate

## future work

see GitHub project : https://github.com/hrkt/awase-server/projects

# How to run

0. edit app-settings.json

see: "app-settings.json" paragraph in this README. 

1. execute server.

```
$ make run
```

2. make HTTP POST request

```
$ curl -X GET localhost:8080/api/event/youreventid"
{"ID":"xxxxxxx", WORK_IN_PROGRESS}
```


# prerequisites

- dep as dependency manager
- linux, *nix like platforms - ("endless" shows error message  on windows platform, at this point of moment)

# app-settings.json

specify "XX" and "XX"

```
{
    "XX": "xxx",
    "XX": [
        "aaa",
        "bbb"
    ]
}
```


# usage

## run in dev mode

```
$ make run
```

## run in release mode

```
$ make release-run
```

## build

```
$ make run
```

## test

```
$ make test
```

## format

```
$ make fmt
```

## build container

```
$ make build-container
```

## run container

```
$ make run-container
```

# License
MIT

# CI

[![CircleCI](https://circleci.com/gh/hrkt/cmd-exec-server.svg?style=svg)](https://circleci.com/gh/hrkt/cmd-exec-server)