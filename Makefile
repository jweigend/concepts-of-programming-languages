# 
# Standard Makefile. Usage: make <target>. Exclude raft package (slow tests!)
# 
PKGS := $(shell go list ./... | grep -v /raft)

.PHONY: test all build test install clean slideshow help coverage

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
	@echo "       make coverage : Executes the tests with coverage and starts the go tool cover"
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
	rm *.out 

all: build test install

coverage:
	go test -coverprofile=coverage.out $(PKGS)
	go tool cover -html=coverage.out

slideshow:
	cd docs; present
