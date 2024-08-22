DOCKERCMD=docker

SERVICE_NAME=multifinance-api

DOCKER_CONTAINER_NAME?=$(SERVICE_NAME)
DOCKER_CONTAINER_IMAGE?=$(SERVICE_NAME):latest
DOCKER_BUILD_ARGS?=
DOCKER_DEBIAN_MIRROR?=http://deb.debian.org/debian

BUILD_DATE?=$(shell date -u +'%Y-%m-%dT00:00:00Z')
BUILD_VERSION?=0.1.0

TOPDIR=$(PWD)
BINARY=$(SERVICE_NAME)

.FORCE:
.PHONY: build
.PHONY: vet
.PHONY: unit-test
.PHONY: generate
.PHONY: depend
.PHONY: docker-build
.PHONY: clean
.PHONY: install
.PHONY: all
.PHONY: .FORCE

build:
	@echo "Executing go build"
	go build -v -buildmode=pie -ldflags "-X main.version=$(BUILD_VERSION)"
	@echo "Binary ready"

vet:
	@echo "Running Go static code analysis with go vet"
	go vet -asmdecl -atomic -bool -buildtags -copylocks -httpresponse -loopclosure -lostcancel -methods -nilfunc -printf -rangeloops -shift -structtag -tests -unreachable -unsafeptr ./...
	@echo "go vet complete"

unit-test:
	@echo "Executing go unit test"
	go test -v -json -count=1 -parallel=4 ./...
	@echo "Unit test done"

generate:
	go generate ./...

depend:
	@echo "Pulling all Go dependencies"
	go mod tidy
	go mod download
	go mod verify
	@echo "You can now run 'make install' to compile all packages"

docker-build:
	$(DOCKERCMD) build -t $(DOCKER_CONTAINER_IMAGE) --build-arg DEBIAN_MIRROR=$(DOCKER_DEBIAN_MIRROR) --build-arg BUILD_ID=$(BUILD_ID) --build-arg BUILD_DATE=$(BUILD_DATE) --build-arg GOPROXY=$(GOPROXY) --build-arg GOSUMDB=$(GOSUMDB) --build-arg BUILD_VERSION=$(BUILD_VERSION) --build-arg GIT_USERNAME=$(GIT_USERNAME) --build-arg GIT_PASSWORD=$(GIT_PASSWORD) $(DOCKER_BUILD_ARGS) .

default: depend

all: depend generate build unit-test

install: depend build

clean:
	rm -f $(BINARY)
	rm -f $(BINARY).exe