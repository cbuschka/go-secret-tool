PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

build:
	cd ${PROJECT_DIR} && \
	go build -ldflags "-extldflags \"-static\"" -o dist/secret-tool ./cmd/...

tidy:
	cd ${PROJECT_DIR} && \
	go mod verify
