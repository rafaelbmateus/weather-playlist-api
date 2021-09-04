.PHONY: build
build: ##@development Build docker image
	docker-compose build

.PHONY: up
up: ##@development Start development environment in background.
	docker-compose up -d

.PHONY: logs
logs: ##@development Follows development logs [service="svc1 svc2..."].
	docker-compose logs -f --tail=100 $(service)

.PHONY: shell
shell: ##@development Start a shell session within the container.
	docker-compose run --rm app /bin/sh

lint_version ?= v1.40-alpine
.PHONY: lint
lint: ##@development Run static analysis code.
	docker run --rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint:$(lint_version) \
		golangci-lint run --timeout 3m

.PHONY: test
test: ##@development Run the tests.
	docker-compose run --rm app go test -v ./...

.PHONY: clean
clean: ##@development Stop development environment.
	docker-compose down -v --remove-orphans
