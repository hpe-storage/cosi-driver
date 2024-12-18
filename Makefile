# © Copyright 2024 Hewlett Packard Enterprise Development LP
IMAGE ?= cosi-driver
ifndef REPO_NAME
    REPO_NAME ?= hpestorage/$(IMAGE)
endif
IMAGE_NAME ?= $(REPO_NAME)
ifdef REGISTRY_NAME
	IMAGE_NAME = $(REGISTRY_NAME)/$(REPO_NAME)
endif
IMAGE_TAGS ?= latest
ARCH ?= amd64
HTTP_PROXY ?=
HTTPS_PROXY ?=
RUN_TESTS ?= true

GOENV = PATH=$$PATH:$(GOPATH)/bin

.PHONY: help
help:
	@echo "Targets:"
	@echo "    clean        - Remove build artifacts."
	@echo "    build        - Compiles the source code."
	@echo "    test         - Run unit tests."
	@echo "    container    - Build cosi driver image and create a local docker image.  Errors are ignored."
	@echo "    push         - Build and push cosi driver image to registry."
	@echo "    all          - Cleans build artifacts, build and run unit tests and push the container image to the registry."

.PHONY: build
build:
	@mkdir -p bin
	@CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -a -ldflags '-extldflags "-static"' -o ./bin/$(IMAGE) ./*.go

.PHONY: container
container:
	docker buildx build --platform=linux/amd64,linux/arm64 --progress=plain \
		$(if $(HTTP_PROXY),--build-arg http_proxy=$(HTTP_PROXY)) \
		$(if $(HTTPS_PROXY),--build-arg https_proxy=$(HTTPS_PROXY)) \
		--build-arg RUN_TESTS=$(RUN_TESTS) \
		--provenance=false $(addprefix -t $(IMAGE_NAME):,$(IMAGE_TAGS)) .

.PHONY: push
push:
	docker buildx build --platform=linux/amd64,linux/arm64 --progress=plain \
		$(if $(HTTP_PROXY),--build-arg http_proxy=$(HTTP_PROXY)) \
		$(if $(HTTPS_PROXY),--build-arg https_proxy=$(HTTPS_PROXY)) \
		--build-arg RUN_TESTS=$(RUN_TESTS) \
		--provenance=false --push $(addprefix -t $(IMAGE_NAME):,$(IMAGE_TAGS)) .

.PHONY: test
test:
	@echo "Running unit tests"
	go test -v -gcflags=all=-l ./...

.PHONY: clean
clean:
	@rm -rf bin

.PHONY: all
all:
	$(MAKE) push RUN_TESTS=$(RUN_TESTS)
