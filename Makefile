PWD := $(shell pwd)
APP := puppy
PKG := github.com/ckeyer/$(APP)

GO := CGO_ENABLED=0 GOBIN=${PWD}/bundles go
HASH := $(shell which sha1sum || which shasum)

OS := $(shell go env GOOS)
ARCH := $(shell go env GOARCH)
VERSION := $(shell cat VERSION.txt)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_AT := $(shell date "+%Y-%m-%dT%H:%M:%SZ%z")
PACKAGE_NAME := $(APP)$(VERSION).$(OS)-$(ARCH)

LD_FLAGS := -X github.com/ckeyer/commons/version.version=$(VERSION) \
 -X github.com/ckeyer/commons/version.gitCommit=$(GIT_COMMIT) \
 -X github.com/ckeyer/commons/version.buildAt=$(BUILD_AT) -w

DEV_IMAGE := ckeyer/dev:go

gorun:
	$(GO) run -ldflags="$(LD_FLAGS)" cli/main.go

local:
	$(GO) install -a -ldflags="$(LD_FLAGS)" .
	make hash

build:
	$(GO) build -a -ldflags="$(LD_FLAGS)" -o bundles/$(APP) main.go
	make hash

hash:
	$(HASH) bundles/$(APP)

test:
	$(GO) test -ldflags="$(LD_FLAGS)" $$(go list ./... |grep -v "vendor")

release: clean local
	mkdir -p bundles/$(PACKAGE_NAME)
	mv bundles/$(APP) bundles/$(PACKAGE_NAME)
	cd bundles ;\
	 echo $(VERSION) > $(PACKAGE_NAME)/release.txt ;\
	 $(HASH) $(PACKAGE_NAME)/$(APP) > $(PACKAGE_NAME)/sha1.txt ;\
	 tar zcvf $(PACKAGE_NAME).tgz $(PACKAGE_NAME);

clean:
	rm -rf bundles/*

dev:
	docker run --rm -it \
	 --name $(APP)-dev \
	 -p 8000:8000 \
	 $(NET) \
	 -v /var/run/docker.sock:/var/run/docker.sock \
	 -v $(PWD)/..:/opt/gopath/src/$(PKG)/.. \
	 -w /opt/gopath/src/$(PKG) \
	 $(DEV_IMAGE) bash
