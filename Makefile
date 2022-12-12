GO=go
export GOPRIVATE=github.com/woodwo/*
TARGET_DIR?=$(PWD)/.build
M=$(shell printf "\033[34;1m>>\033[0m")


.PHONY: build
build: 
	$(info $(M) building service...)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o $(TARGET_DIR)/client ./cmd/*.go

.PHONY: run
run: build 
	$(TARGET_DIR)/client

.PHONY: tidy
tidy:
	go mod tidy
%:
	@:

