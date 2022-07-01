# Variable
PROJECTNAME=$(shell basename "$(PWD)")
WINDOWS=$(PROJECTNAME)_windows_amd64.exe
LINUX=$(PROJECTNAME)_linux_amd64
DARWIN=$(PROJECTNAME)_darwin_amd64
OBJECTS=$(WINDOWS) $(LINUX) $(DARWIN)
VERSION=$(shell git describe --tags $(git rev-list --tags --max-count=1))
COMMIT=$(shell git rev-parse --short HEAD)
LDFLAGS=-ldflags="-X 'main.version=${VERSION}' -X 'main.commit=${COMMIT}'"
# option
CGO_ENABLED_FLAG:=CGO_ENABLED=1

echo:
	echo $(detected_OS)
	echo $(WINDOWS)
	echo $(LINUX)
	echo $(DARWIN)
	echo $(CFLAGS)
	echo $(MAKE)

$(WINDOWS):
# 	$(CGO_ENABLED_FLAG) GOOS=windows GOARCH=amd64 $(GO) build -v -o $(WINDOWS) main.go
# 	@echo "  >  $(CGO_ENABLED_FLAG) GOOS=windows GOARCH=amd64 $(GO) build -v -o $(WINDOWS) main.go"
	$(CGO_ENABLED_FLAG) GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -v -o $(WINDOWS) main.go
	@echo "  >  $(CGO_ENABLED_FLAG) GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(WINDOWS) main.go"

$(LINUX):
# 	$(CGO_ENABLED_FLAG) GOOS=linux GOARCH=amd64 $(GO) build -v -o $(LINUX) main.go
# 	@echo "  >  $(CGO_ENABLED_FLAG) GOOS=windows GOARCH=amd64 $(GO) build -v -o $(WINDOWS) main.go"
	$(CGO_ENABLED_FLAG) GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -v -o $(LINUX) main.go
	@echo "  >  $(CGO_ENABLED_FLAG) GOOS=windows GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(LINUX) main.go"

$(DARWIN):
# 	$(CGO_ENABLED_FLAG) GOOS=darwin GOARCH=amd64 $(GO) build -v -o $(DARWIN) main.go
# 	@echo "  >  $(CGO_ENABLED_FLAG) GOOS=darwin GOARCH=amd64 $(GO) build -v -o $(DARWIN) main.go"
	$(CGO_ENABLED_FLAG) GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -v -o $(DARWIN) main.go
	@echo "  >  $(CGO_ENABLED_FLAG) GOOS=darwin GOARCH=amd64 $(GO) build $(LDFLAGS) -o $(DARWIN) main.go"

# Check os
UNAME := $(shell uname)
ifeq ($(OS),Windows_NT) # is Windows_NT on XP, 2000, 7, Vista, 10...
	detected_OS := Windows
else
	detected_OS := $(shell sh -c 'uname 2>/dev/null || echo Unknown')
endif

ifeq '$(findstring ;,$(PATH))' ';'
    detected_OS := Windows
else
    detected_OS := $(shell uname 2>/dev/null || echo Unknown)
    detected_OS := $(patsubst CYGWIN%,Cygwin,$(detected_OS))
    detected_OS := $(patsubst MSYS%,MSYS,$(detected_OS))
    detected_OS := $(patsubst MINGW%,MSYS,$(detected_OS))
endif

ifeq ($(detected_OS),Windows)
	CFLAGS += $(MAKE) $(WINDOWS)
endif
ifeq ($(detected_OS),Darwin)        # Mac OS X
# 	CFLAGS += $(MAKE) $(WINDOWS)
	CFLAGS += $(MAKE) $(DARWIN)
# 	CFLAGS += $(MAKE) $(LINUX)
endif
ifeq ($(detected_OS),Linux)
	CFLAGS += $(MAKE) $(LINUX)
endif
ifeq ($(detected_OS),GNU)           # Debian GNU Hurd
	CFLAGS +=-D GNU_HURD
endif
ifeq ($(detected_OS),GNU/kFreeBSD)  # Debian kFreeBSD
	CFLAGS +=-D GNU_kFreeBSD
endif
ifeq ($(detected_OS),FreeBSD)
	CFLAGS +=-D FreeBSD
endif
ifeq ($(detected_OS),NetBSD)
	CFLAGS +=-D NetBSD
endif
ifeq ($(detected_OS),DragonFly)
	CFLAGS +=-D DragonFly
endif
ifeq ($(detected_OS),Haiku)
	CFLAGS +=-D Haiku
endif

# Go build rules.
GO ?= go
# Go related variables.
GOBASE=$(shell pwd)
# GOPATH="/Users/blakewu/go"
GOBIN=$(GOBASE)
GOFILES=$(wildcard *.go)
ifndef $(GOPATH)
	GOPATH=$(shell go env GOPATH)
	export GOPATH
endif

# Redirect error output to a file, so we can show it in development mode.
STDERR=stderr.txt

# PID file will keep the process id of the server
PID=.pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: go-get

## start: Start in development mode. Auto-starts when code changes.
start:
	bash -c "trap 'make stop' EXIT; $(MAKE) compile start-server watch run='make compile start-server'"

## stop: Stop development mode.
stop: stop-server

start-server: stop-server
	@echo "  >  $(PROJECTNAME) is available at $(ADDR)"
	@-$(GOBIN)/$(PROJECTNAME) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop-server:
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)

## watch: Run given command when code changes. e.g; make watch run="echo 'hey'"
watch:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) yolo -i . -e vendor -e bin -c "$(run)" -a localhost:8088

restart-server: stop-server start-server

## compile: Compile the binary.
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
@cat $(STDERR) | sed -e '1s/.*/\nError:\n/'  | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

## exec: Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
exec:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)

## clean: Clean build files. Runs `go clean` internally.
clean:
	@-$(MAKE) go-clean
	@-rm $(WINDOWS) $(LINUX) $(DARWIN)

go-compile: clean go-build

go-build:
	@echo "  >  Building binary..."
	$(CFLAGS)
	echo "  >  $(CFLAGS)"
	@echo "  >  Building binary end"

go-generate:
	@echo "  >  Generating dependency files..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO) generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO) get .

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO) install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(GO) clean

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo