# time-tracker

## Prerequisites

### devcontainers/cli

```bash
npm install -g @devcontainers/cli
```

## Requirements

*   docker 26.0.0
*   docker-compose 1.29.2
*   devcontainers/cli (optional)

## Build

### Docker

You can build a project manually with `docker`. You have to build an image and run a containers with commands:

```bash
docker build \
	-t base-time-tracker:1.0 \
	-f docker/base.build.Dockerfile \
	.

docker build \
	-t time-tracker:1.0 \
	-f docker/timetracker.Dockerfile \
	.

docker run \
    --rm \
    -p 0.0.0.0:8080:8080 \
    time-tracker:1.0 \
    .
```

### Dev Container Cli

Also, you can build a project with [devcontainers](https://containers.dev/) in an easy and convenient way.

Your IDE or code editor can run and attach to devcontainer.

You can use devcontainers/cli to set up environment and build the project manually via bash:

```bash
docker build \
	-t base-time-tracker:1.0 \
	-f docker/base.build.Dockerfile \
	.

devcontainer up --workspace-folder .

devcontainer exec --workspace-folder . go run /go/cmd/main.go
```

## Run

You can run application by building binary:

```bash
go build -o ./bin/timetrackerbin ./cmd/main.go

go run ./bin/timetrackerbin
```

Also you can run dockerized application via docker compose:
```bash
docker compose -f deploy/docker-compose.yml up --build
```
