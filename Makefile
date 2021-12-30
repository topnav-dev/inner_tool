# Linux、Cygwin、MSYS、Windows、FreeBSD、NetBSD、Solaris、Darwin、OpenBSD、AIX、HP-UX
# ifeq '$(findstring ;,$(PATH))' ';'
#     UNAME = Windows
# else
#     UNAME := $(shell uname 2>/dev/null || echo Unknown)
#     UNAME := $(patsubst CYGWIN%,Cygwin,$(UNAME))
#     UNAME := $(patsubst MSYS%,MSYS,$(UNAME))
#     UNAME := $(patsubst MINGW%,MSYS,$(UNAME))
# endif

# MAIN_FILE=main.go version.go

EXECUTABLE=cmd
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --dirty | sed 's/-g[a-z0-9]\{7\}//')
COMMIT=$(shell git rev-parse --short HEAD)

LDFLAGS=-ldflags="-w -s \
-X 'main.versionString=${VERSION}' \
-X 'main.commitString=${COMMIT}'"

OBJECTS=$(WINDOWS) $(LINUX) $(DARWIN)

$(WINDOWS):
	GOOS=windows GOARCH=amd64
	@go build -v -o $(WINDOWS) $(LDFLAGS) .

$(LINUX):
	GOOS=linux GOARCH=amd64
	@go build -v -o $(LINUX) $(LDFLAGS) .

$(DARWIN):
	GOOS=darwin GOARCH=amd64
	@go build -v -o $(DARWIN) $(LDFLAGS) .

build: $(OBJECTS) ## Build binaries
	@echo versionString: $(VERSION)
	@echo commitString: $(COMMIT)

move:
	mv $(OBJECTS) ../

run: build move

clean: ## Remove previous build
	@go clean
	rm -f ../$(DARWIN)
	rm -f ../$(LINUX)
	rm -f ../$(WINDOWS)

test: ## Run unit tests
	./scripts/test_unit.sh

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run clean