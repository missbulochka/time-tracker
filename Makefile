# You can run docker-containers with service with command
run-app: build-base
	docker-compose -f deploy/docker-compose.yml up

stop-app:
	docker-compose -f deploy/docker-compose.yml down

# You can build base with command
build-base:
	docker build \
		-t base-time-tracker:1.0 \
		-f docker/base.build.Dockerfile \
		.

# You can build image with service with command
build-app: build-base
	docker build \
		-t time-tracker:1.0 \
		-f docker/timetracker.Dockerfile \
		.

# You can up devcontainer with command
up-dev: build-base
	devcontainer up --workspace-folder .

# You can run your app with devcontainer
run-dev: up-dev
	devcontainer exec --workspace-folder . go run /usr/src/timetracker/cmd/main.go

# You can down dev-container with
down-dev:
	docker stop dev_time-tracker time-tracker-psql
	docker rm dev_time-tracker time-tracker-psql
