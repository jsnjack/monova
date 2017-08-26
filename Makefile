BINARY:=dahary
DAHARY:=$(shell dahary -v dot 2> /dev/null)
VERSION:=0.0.0

version:
ifdef DAHARY
	VERSION:=$(shell dahary)
else
	$(info !Install dahary to automatically update version)
endif

build: version
	go build -ldflags="-X main.version=${VERSION}"

run: build
	./${BINARY}
