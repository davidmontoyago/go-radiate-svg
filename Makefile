# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GO111MODULE=off

all: test build

build:
	$(GOBUILD) ./

test:
	$(GOTEST) ./

clean:
	$(GOCLEAN)

fmt:
	$(GOCMD) fmt ./...

run:
	go run main.go > radiator.html
