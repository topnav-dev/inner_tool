#Linux、Cygwin、MSYS、Windows、FreeBSD、NetBSD、Solaris、Darwin、OpenBSD、AIX、HP-UX
# ifeq '$(findstring ;,$(PATH))' ';'
#     UNAME = Windows
# else
#     UNAME := $(shell uname 2>/dev/null || echo Unknown)
#     UNAME := $(patsubst CYGWIN%,Cygwin,$(UNAME))
#     UNAME := $(patsubst MSYS%,MSYS,$(UNAME))
#     UNAME := $(patsubst MINGW%,MSYS,$(UNAME))
# endif

MAIN_FILE=main.go

EXECUTABLE=conv
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

OBJECTS=$(WINDOWS) $(LINUX) $(DARWIN)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" $(MAIN_FILE)

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" $(MAIN_FILE)

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" $(MAIN_FILE)

build: $(OBJECTS) ## Build binaries
	@echo version: $(VERSION)

run: build

clean: ## Remove previous build
	rm -f ../$(OBJECTS)

test: ## Run unit tests
	./scripts/test_unit.sh

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run clean