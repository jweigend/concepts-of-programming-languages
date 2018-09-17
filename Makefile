# 
# Standard Makefile. Usage: make <target>.
# 
PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test all build test install clean slideshow help

help:
	clear
	@echo "---------------------------------------------------------------------------------------------------------"
	@echo "Usage: make <target>"
	@echo "---------------------------------------------------------------------------------------------------------"
	@echo "Valid targets are:"
	@echo "       make all : Runs build, test, install."
	@echo "       make build : Builds all packages."
	@echo "       make test : Runs all tests."
	@echo "       make install : Installs all packages."
	@echo "       make clean : Clean up and clears caches."
	@echo "       make slideshow : Starts a golang present slideshow on port 3999. Blocks until CTRL-C ist pressed. "
	@echo "       make help : This info.                                                                            "
	@echo "---------------------------------------------------------------------------------------------------------"

test:
	go test $(PKGS)

build:
	go build $(PKGS)

install: 
	go install $(PKGS)

clean:
	go clean -testcache -cache $(PKGS)

all: build test install

slideshow:
	cd doc; present
