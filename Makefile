# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go test -race -vet=off ./...
	go mod verify


# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build: build the cmd/api application
.PHONY: build
build:
	go mod verify
	go build -ldflags='-s' -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api

## run: run the cmd/api application
.PHONY: run
run: tidy build
	./bin/api


# ==================================================================================== #
# SQL MIGRATIONS
# ==================================================================================== #

## migrations/new name=$1: create a new database migration
.PHONY: migrations/new
migrations/new:
	go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -seq -ext=.sql -dir=./assets/migrations ${name}

## migrations/up: apply all up database migrations
.PHONY: migrations/up
migrations/up:
	go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="mysql://${DB_DSN}" up

## migrations/down: apply all down database migrations
.PHONY: migrations/down
migrations/down:
	go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="mysql://${DB_DSN}" down

## migrations/goto version=$1: migrate to a specific version number
.PHONY: migrations/goto
migrations/goto:
	go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="mysql://${DB_DSN}" goto ${version}

## migrations/force version=$1: force database migration
.PHONY: migrations/force
migrations/force:
	go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="mysql://${DB_DSN}" force ${version}

## migrations/version: print the current in-use migration version
.PHONY: migrations/version
migrations/version:
	go run -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="mysql://${DB_DSN}" version

