all: test install

build:
	go build

install:
	go install

lint:
	golint .

test:
	go test -v
