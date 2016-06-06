# variable definitions
NAME := xget
DESC := a multi-connections file downloader
PREFIX ?= usr/local
VERSION := $(shell git describe --tags --always --dirty)
GOVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDDATE := $(shell date -u +"%B %d, %Y")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>")
PKG_RELEASE ?= 1
PROJECT_URL := "https://github.com/miraclew/$(NAME)"
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.buildTime=$(BUILDTIME)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(GOVERSION)'


# development tasks
test:
    go test $$(go list ./... | grep -v /vendor/ | grep -v /cmd/)

coverage:
    @-go test -v -coverprofile=cover.out $$(go list ./... | grep -v /vendor/ | grep -v /cmd/)
    @-go tool cover -html=cover.out -o cover.html

benchmark:
    @echo "Running tests..."
    @go test -bench=. $$(go list ./... | grep -v /vendor/ | grep -v /cmd/)