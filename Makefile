SHELL := /bin/bash

TARGET := $(shell echo $${PWD\#\#*/})
OUT := "bin/"
.DEFAULT_GOAL: $(TARGET)

VERSION := 0.0.1
#BUILD := `git rev-parse HEAD`

#LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all build clean install uninstal fmt simplify check run

all: build

$(TARGET): $(SRC)
	@cd src && \
	go build $(LDFLAGS) -o ../${OUT}${TARGET}

build: $(TARGET)
	@true

clean:
	@rm -f ${OUT}${TARGET}

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

gen:
	@cd src && \
	go generate
