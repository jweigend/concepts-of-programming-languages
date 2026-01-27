# 
# Standard Makefile. Usage: make <target>. Exclude raft package (slow tests!)
# 
PKGS := $(shell go list ./... | grep -v /raft)

.PHONY: test all build test install clean slideshow help coverage pdfs

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
	@echo "       make pdfs : Generates PDF files from all .slide files (requires Chrome). "
	@echo "       make help : This info.                                                                            "
	@echo "---------------------------------------------------------------------------------------------------------"

test:
	go test $(PKGS)

build:
	go build $(PKGS)

install: 
	go install $(PKGS)

clean:
	go clean -testcache -cache
	-rm -f *.out 

all: build test install

coverage:
	go test -coverprofile=coverage.out $(PKGS)
	go tool cover -html=coverage.out

# GOBIN determines where 'go install' places binaries
GOBIN := $(shell go env GOBIN)
ifeq ($(GOBIN),)
	GOBIN := $(shell go env GOPATH)/bin
endif

# Check if present is installed, install if missing, then run from docs directory
slideshow:
	@test -x "$(GOBIN)/present" || { \
		echo "Installing 'present' tool to $(GOBIN)..."; \
		go install golang.org/x/tools/cmd/present@latest; \
	}
	cd docs && $(GOBIN)/present -notes

# Generate PDFs from .slide files using Selenium + Chrome
# Starts present server from docs directory, generates PDFs, then stops server
pdfs:
	@echo "Generating PDFs from slide files..."
	@test -x "$(GOBIN)/present" || { \
		echo "Installing 'present' tool to $(GOBIN)..."; \
		go install golang.org/x/tools/cmd/present@latest; \
	}
	@cd docs && $(GOBIN)/present -notes & \
	sleep 2; \
	python3 scripts/generate_pdfs.py --wait 3; \
	pkill -f "present -notes" 2>/dev/null || true
